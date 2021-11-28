package com.example.javaspring;

import com.example.javaspring.model.Task;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import com.example.javaspring.service.ITaskService;

import java.util.List;

@RestController
public class TaskController {

//    @Autowired
    private ITaskService taskService;

    public TaskController(ITaskService taskService) {
        this.taskService = taskService;
    }

    @GetMapping("/tasks/{id}")
    public Task findById(@PathVariable Long id) {
        return taskService.findById(id);
    }

    @GetMapping("/tasks")
    public List<Task> findAll() {
        return taskService.findAll();
    }

    @PostMapping("/tasks")
    public Task createTask(@RequestBody Task newTask) {
        return taskService.createTask(newTask);
    }



    @PutMapping("/tasks/{id}")
    public Task updateTask(@RequestBody Task newTask, @PathVariable Long id) {
        return taskService.updateTask(newTask, id);
    }

    @DeleteMapping("tasks/{id}")
    public void deleteTask(@PathVariable Long id) {
        taskService.deleteTask(id);
    }
}
