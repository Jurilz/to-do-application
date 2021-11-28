import * as React from 'react'

import {TodoInterface} from "./todo-interface";

export interface TodoItemInterface {
    handleTodoUpdate: (event: React.ChangeEvent<HTMLInputElement>, id: string) => void;
    handleTodoRemove: (id: string) => void;
    handleTodoComplete: (id: string) => void;
    todo: TodoInterface;
}
