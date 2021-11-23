package com.example.javaspring.service;

import com.example.javaspring.model.Task;
import com.example.javaspring.repository.TaskRepository;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import javax.persistence.EntityNotFoundException;
import java.util.List;

@Service
public class TaskService implements ITaskService{

    @Autowired
    private TaskRepository taskRepository;

    @Override
    public List<Task> findAll() {
        return taskRepository.findAll();
    }

    @Override
    public Task findById(Long taskId) {
        return taskRepository.findById(taskId)
                .orElseThrow(EntityNotFoundException::new);
    }

    @Override
    public Task createTask(Task newTask) {
        return taskRepository.save(newTask);
    }

    @Override
    public Task updateTask(Task newTask, Long taskId) {
        return taskRepository.findById(taskId)
                .map(task -> {
                    task.setLabel(newTask.getLabel());
                    task.setDone(newTask.getDone());
                    task.setDate(newTask.getDate());
                    return taskRepository.save(task);
                })
                .orElseGet(() -> {
                    newTask.setId(taskId);
                    return taskRepository.save(newTask);
                });
    }

    @Override
    public void deleteTask(Long taskId) {
        taskRepository.deleteById(taskId);
    }


}
