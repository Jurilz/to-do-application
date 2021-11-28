import * as React from "react";

import ToDoItem from "./ToDoItem";
import {TodoListInterface} from "../interfaces/todo-list-interace";

const ToDoList = (props: TodoListInterface) => {
    return (
        <div className="todo-list">
            <ul>
                { props.todos.map( (todo) => (
                    <li key={todo.id}>
                        <ToDoItem
                            todo={todo}
                            handleToDoUpdate={props.handleTodoUpdate}
                            handleTodoRemove={props.handleTodoRemove}
                            handleTodoComplete={props.handleTodoComplete}
                            />
                    </li>
                ) )}
            </ul>

        </div>
    )
};
export default ToDoList;
