package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/VG-Tech-Dojo/treasure2018/mid/NAKKA-K/cook-do/model"
	"launchpad.net/xmlpath"

	"github.com/jmoiron/sqlx"
)

// Todo はTodoへのリクエストに関する制御をします
type Todo struct {
	DB *sqlx.DB
}

// Get はDBからユーザを取得して結果を返します
func (t *Todo) Get(w http.ResponseWriter, r *http.Request) error {
	todos, err := model.TodosAll(t.DB)
	if err != nil {
		return err
	}
	return JSON(w, 200, todos)
}

// Put はRequestされたTodoを更新します
func (t *Todo) Put(w http.ResponseWriter, r *http.Request) error {
	var todo model.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		return err
	}

	if err := TXHandler(t.DB, func(tx *sqlx.Tx) error {
		result, err := todo.Update(tx)
		if err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		todo.ID, err = result.LastInsertId()
		return err
	}); err != nil {
		return err
	}

	return JSON(w, http.StatusOK, todo)
}

// Post はタスクをDBに追加します
// todoをJSONとして受け取ることを想定しています。
func (t *Todo) Post(w http.ResponseWriter, r *http.Request) error {
	var todo model.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		return err
	}

	if err := TXHandler(t.DB, func(tx *sqlx.Tx) error {
		result, err := todo.Insert(tx)
		if err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		todo.ID, err = result.LastInsertId()
		return err
	}); err != nil {
		return err
	}

	return JSON(w, http.StatusCreated, todo)
}

// Delete はRequestされたTodoをdeleteします
func (t *Todo) Delete(w http.ResponseWriter, r *http.Request) error {
	var todo model.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		return err
	}

	if err := TXHandler(t.DB, func(tx *sqlx.Tx) error {
		_, err := todo.Delete(tx)
		if err != nil {
			return err
		}
		return tx.Commit()
	}); err != nil {
		return err
	}

	return JSON(w, http.StatusOK, todo)
}

// Toggle はRequestされたTodoのcompletedをトグルします
func (t *Todo) Toggle(w http.ResponseWriter, r *http.Request) error {
	var todo model.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		return err
	}

	if err := TXHandler(t.DB, func(tx *sqlx.Tx) error {
		result, err := todo.Toggle(tx)
		if err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		todo.ID, err = result.LastInsertId()
		return err
	}); err != nil {
		return err
	}

	return JSON(w, http.StatusOK, todo)
}

// DeleteAll は全てのTodoをDeleteします
func (t *Todo) DeleteAll(w http.ResponseWriter, r *http.Request) error {
	if err := TXHandler(t.DB, func(tx *sqlx.Tx) error {
		_, err := model.TodosDeleteAll(tx)
		if err != nil {
			return err
		}
		return tx.Commit()
	}); err != nil {
		return err
	}
	var res interface{}
	return JSON(w, http.StatusOK, res)
}

// Search はidまたはtitleをクエリパラメータとして受け取って、検索結果のTodo配列を返します
func (t *Todo) Search(w http.ResponseWriter, r *http.Request) error {
	if idStr := r.URL.Query().Get("id"); idStr != "" {
		todo, err := t.searchID(w, idStr)
		if err != nil {
			return err
		} else if todo == nil {
			return returnEmptyJSON(w)
		}
		return JSON(w, 200, todo)
	}

	title := r.URL.Query().Get("title")
	if title == "" {
		return returnEmptyJSON(w)
	}

	searched, err := t.searchTitle(w, title)
	if err != nil {
		return err
	} else if searched == nil {
		return returnEmptyJSON(w)
	}

	return JSON(w, 200, searched)
}

func (t *Todo) searchTitle(w http.ResponseWriter, title string) ([]model.Todo, error) {
	todos, err := model.TodosAll(t.DB)
	if err != nil {
		return nil, err
	}

	seached := make([]model.Todo, 0, 5)
	for _, v := range todos {
		if strings.Contains(v.Title, title) {
			seached = append(seached, v)
		}
	}

	return seached, nil
}

func (t *Todo) searchID(w http.ResponseWriter, idStr string) (*model.Todo, error) {
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return nil, err
	}
	todo, err := model.TodoOne(t.DB, id)
	if err != nil {
		return nil, nil
	}
	return todo, nil
}

func returnEmptyJSON(w http.ResponseWriter) error {
	return JSON(w, 200, make([]model.Todo, 0))
}

// Scraping はクックパッドレシピサイトのURLを受け取ってスクレイピングする
// スクレイピングした調理手順をTODOに落とし込んで、配列として返す
func (t *Todo) Scraping(w http.ResponseWriter, r *http.Request) error {
	var url string
	if err := json.NewDecoder(r.Body).Decode(&url); err != nil {
		return err
	}

	todos, err := scrapingCookpad(url)
	if err != nil {
		return err
	}

	// TODO: レスポンスとして400番を返しながら、JSONのbodyにエラーメッセージを返したい
	if todos == nil {
		return JSON(w, 400, fmt.Errorf("Can't access to url"))
	}

	if err := TXHandler(t.DB, func(tx *sqlx.Tx) error {
		err := todos.InsertAll(tx)
		if err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		return err
	}); err != nil {
		return err
	}

	return JSON(w, http.StatusCreated, *todos)
}

func scrapingCookpad(url string) (*model.Todos, error) {
	timeout := time.Duration(5 * time.Second)
	client := http.Client{Timeout: timeout}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ParseCookpadRecipe(&resp.Body)
}

// ParseCookpadRecipe は*io.ReadCloser型を受け取ってパースします。
// パースした結果を*model.Todosに格納し返却します。
func ParseCookpadRecipe(body *io.ReadCloser) (*model.Todos, error) {
	root, err := xmlpath.ParseHTML(*body)
	if err != nil {
		return nil, err
	}

	title := extractTitle(root)
	todos := parseStepToTodo(root, title)
	return todos, nil
}

func extractTitle(root *xmlpath.Node) string {
	titlePath := xmlpath.MustCompile(`//div[@id="recipe-title"]/h1`)
	iter := titlePath.Iter(root)
	for iter.Next() {
		return strings.TrimSpace(iter.Node().String())
	}
	return ""
}

// parseStepToTodo はクックパッドの調理手順一覧をパースして、model.Todosに格納する
func parseStepToTodo(root *xmlpath.Node, title string) *model.Todos {
	path := xmlpath.MustCompile(`//*[@id="steps"]/div/dl/dd/p`) // クックパッドの調理手順一覧
	iter := path.Iter(root)
	var todos model.Todos
	for i := 1; iter.Next(); i++ {
		recipeHead := "[" + title + "]: " + strconv.Itoa(i) + ". "
		todo := &model.Todo{
			Title:     recipeHead + strings.TrimSpace(iter.Node().String()),
			Completed: false,
		}
		todos = append(todos, *todo)
	}
	return &todos
}
