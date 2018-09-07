export function rootReducer(
    state = { isLoading: false, tasks: [], token: null },
    action
) {
    switch (action.type) {
        case "SET_TASKS": // csrf token
            return { ...state, tasks: action.tasks };
        case "ADD_TASK":
            return {
                ...state,
                tasks: [...state.tasks, action.task]
            };
        case "ADD_TASKS":
            return {
                ...state,
                tasks: [...state.tasks, ...action.tasks]
            };
        case "LOADING":
            return { ...state, isLoading: true };
        case "LOADED":
            return { ...state, isLoading: false };
        case "SET_TOKEN":
            return { ...state, token: action.token };
    }
    return state;
}
