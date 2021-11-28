export interface TodoInterface {
    todos: TodoInterface[];
    handleTodoCreate: (todo: TodoInterface) => void;
}