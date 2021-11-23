package com.example.javaspring.service;

import com.example.javaspring.model.Task;

import java.util.List;

public interface ITaskService {

    List<Task> findAll();

    Task findById(Long taskId);

    Task createTask(Task task);

    Task updateTask(Task task, Long taskId);

    void deleteTask(Long taskId);

}
