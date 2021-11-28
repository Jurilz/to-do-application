import * as React from "react";
import { render } from "react-dom";

import ToDoForm from './components/TodoForm';
import ToDoList from './components/ToDoList';

import {TodoInterface} from "./interfaces/todo-interface";

const App: React.FC = () => {
    const [todos, setTodos] = React.useState<TodoInterface[]>([]);

    function handleTodoUpdate(event: React.ChangeEvent<HTMLInputElement>, id: string) {
        const newTodoState: TodoInterface[] = [...todos];

        newTodoState.find( (todo: TodoInterface) => todo.id === id)
            .label = event.target.value;

        setTodos(newTodoState)
    }

    function handleTodoComplete(id: string) {
        const newTodoState: TodoInterface[] = [...todos];

        newTodoState.find( (todo: TodoInterface) => todo.id === id)
            .done = !newTodoState.find( (todo: TodoInterface) => todo.id === id)
            .done;

        setTodos(newTodoState)
    }

    return (
        <div className="App">
            <React.Fragment>
                <h2>My ToDo APP</h2>
                <ToDoForm
        todos={todos}
        handleTodoCreate={handleTodoCreate}
    />
    <ToDoList
        todos={todos}
        handleTodoUpdate={handleTodoUpdate}
        handleTodoRemove={handleTodoRemove}
        handleTodoComplete={handleTodoComplete}

    />

    </React.Fragment>
    </div>
    );
}
export default App;
