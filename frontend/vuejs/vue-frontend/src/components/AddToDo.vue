<template>
    <form @submit="addToDo">
        <input type="text" v-model="title" name="title">
        <button type="submit">Add</button>
    </form>
</template>

<script>
import { v4 as uuidv4 } from 'uuid';

    export default {
        name: "AddToDo",

        // two-way binding of data
        data() {
            return {
                title: ''
            }
        },

        methods: {
            //event is the submit event
            addToDo(event) {
                // default is submitting to a file
                event.preventDefault();

                const newToDo = {
                    id: uuidv4(),
                    title: this.title,
                    completed: false
                };

                this.$emit('add-todo', newToDo);
                this.title = '';
            }
        }
    }
</script>

<style scoped>

</style>
