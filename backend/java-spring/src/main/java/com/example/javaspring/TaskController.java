package com.example.javaspring;

import com.example.javaspring.model.Task;
import com.example.javaspring.repository.TaskSpecificationBuilder;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.jpa.domain.Specification;
import org.springframework.web.bind.annotation.*;

import com.example.javaspring.service.ITaskService;

import java.util.List;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

@RestController
@RequestMapping(path = "/tasks")
public class TaskController {

    @Autowired
    private ITaskService taskService;

    public TaskController(ITaskService taskService) {
        this.taskService = taskService;
    }

    @GetMapping("/{id}")
    public Task findById(@PathVariable Long id) {
        return taskService.findById(id);
    }

    @GetMapping
    public List<Task> findAll(@RequestParam(value = "label") String label) {
        if (label.isEmpty()) {
            return taskService.findAll();
        } else {
            return taskService.findTaskByLabel(label);
        }
    }

    @GetMapping("/filter")
    public List<Task> findByCriteria(@RequestParam(value = "search") String search ) {
        TaskSpecificationBuilder builder = new TaskSpecificationBuilder();
        Pattern pattern = Pattern.compile("(\\w+?)(:|<|>)(\\w+?),");
        Matcher matcher = pattern.matcher(search + ",");
        while (matcher.find()) {
            builder.with(matcher.group(1), matcher.group(2), matcher.group(3));
        }
        Specification<Task> spec = builder.build();
        return taskService.findBySearchCriteria(spec);
    }


    @PostMapping
    public Task createTask(@RequestBody Task newTask) {
        return taskService.createTask(newTask);
    }

    @PutMapping("/{id}")
    public Task updateTask(Task newTask, @PathVariable Long id) {
        return taskService.updateTask(newTask, id);
    }

    @DeleteMapping("/{id}")
    public void deleteTask(@PathVariable Long id) {
        taskService.deleteTask(id);
    }
}
