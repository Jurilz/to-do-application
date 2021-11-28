import * as React from "react";
import shortid from "shortid";

import { TodoInterface } from "../interfaces/todo-interface";
import { TodoFormInterface } from "../interfaces/todo-form-interface";

const ToDoForm = (props: TodoFormInterface) => {

    const inputRef = React.useRef<HTMLInputElement>(null);

    const [values, setValues] = React.useState("");

    function handleInputChange(event: React.ChangeEvent<HTMLInputElement>) {
        setValues(event.target.value);
    }

    function handleInputEnter(event: React.KeyBoardEvent) {
        if (event.key === "Enter") {
            const newTodo: TodoInterface = {
                id: shortid.generate(),
                label: values,
                done: false,
                date: ""
            };

            props.handleTodoCreate(newTodo)

            if (inputRef && inputRef.current) {
                inputRef.current.value = "";
            }
        }
    }

    return (
        <div className="todo-form">
            <input
                ref={inputRef} type="text" placeholder="Enter todo"
                onchange={ event => handleInputChange(event)}
                onkeypress={ event => handleInputEnter(event)}
            />
        </div>
    )
};

export default ToDoForm;

