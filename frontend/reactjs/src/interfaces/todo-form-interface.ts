import { TodoInterface } from "./todo-interface";

export interface TodoFormInterface {
    todos: TodoFormInterface[];
    handleTodoCreate: (todo: TodoInterface) => void;
}
