import { DOMAIN } from "./domain";

export const setTasks = tasks => ({ type: "SET_TASKS", tasks });
export const addTask = task => ({ type: "ADD_TASK", task });
export const addTasks = tasks => ({ type: "ADD_TASKS", tasks });

export const loading = () => ({ type: "LOADING" });
export const loaded = () => ({ type: "LOADED" });

export const setToken = token => ({ type: "SET_TOKEN", token });
export const fetchToken = () => dispatch => {
    fetch(DOMAIN + "/token", { credentials: "same-origin" })
        .then(x => x.json())
        .then(json => {
            if (json === null) {
                return;
            }
            dispatch(setToken(json.token));
        })
        .catch(err => {
            // eslint-disable-next-line
            console.error("fetch error", err);
        });
};

export const loadTodos = () => dispatch => {
    dispatch(loading());
    fetch(DOMAIN + "/api/todos", {
        timeout: 3000,
        mode: "cors",
        method: "GET"
    })
        .then(res => res.json())
        .then(json => {
            if (json !== null) {
                dispatch(setTasks(json));
            }
            dispatch(loaded());
        })
        .catch(dispatch(loaded()));
};

export const postToggle = (toggledTask, tasks) => dispatch => {
    fetch(DOMAIN + "/api/todos/toggle", {
        mode: "cors",
        method: "PUT",
        headers: {
            Accept: "application/json",
            "Content-Type": "application/json"
        },
        body: JSON.stringify(toggledTask)
    })
        .then(res => res.json())
        .then(() => {
            dispatch(
                setTasks(
                    tasks.map(task => {
                        if (task !== toggledTask) {
                            return task;
                        }
                        return Object.assign({}, task, {
                            completed: !task.completed
                        });
                    })
                )
            );
        });
};

export const deleteTask = (task, tasks) => dispatch => {
    return fetch(DOMAIN + "/api/todos", {
        mode: "cors",
        method: "DELETE",
        headers: {
            Accept: "application/json",
            "Content-Type": "application/json"
        },
        body: JSON.stringify(task)
    }).then(() => {
        dispatch(
            setTasks(
                tasks.filter(candidate => {
                    return candidate !== task;
                })
            )
        );
    });
};

export const deleteAllTask = () => dispatch => {
    return fetch(DOMAIN + "/api/todos/all", {
        mode: "cors",
        method: "DELETE",
        headers: {
            Accept: "application/json",
            "Content-Type": "application/json"
        }
    }).then(() => {
        dispatch(setTasks([]));
    });
};

export const postTask = title => dispatch => {
    const task = { title: title, completed: false };

    return fetch(DOMAIN + "/api/todos", {
        mode: "cors",
        method: "POST",
        headers: {
            Accept: "application/json",
            "Content-Type": "application/json"
        },
        body: JSON.stringify(task)
    })
        .then(res => {
            if (res.status === 201) {
                return res;
            }

            throw new Error(res.statusText);
        })
        .then(res => res.json())
        .then(json => dispatch(addTask(json)))
        .catch(err => {
            // eslint-disable-next-line
            console.error("post todo error: ", err);
        });
};

export const postScrapingURL = url => dispatch => {
    return fetch(DOMAIN + "/api/recipe/scraping", {
        mode: "cors",
        method: "POST",
        headers: {
            Accept: "application/json",
            "Content-Type": "application/json"
        },
        body: JSON.stringify(url)
    })
        .then(res => res.json())
        .then(json => dispatch(addTasks(json)))
        .catch(err => {
            // eslint-disable-next-line
            console.error("post recipe url error: ", err);
        });
};
