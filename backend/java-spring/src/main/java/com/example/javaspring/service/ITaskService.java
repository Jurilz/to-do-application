package com.example.javaspring.service;

import com.example.javaspring.model.Task;
import com.example.javaspring.repository.TaskSpecification;
import org.springframework.data.jpa.domain.Specification;
import sun.security.krb5.internal.Ticket;

import java.util.List;

public interface ITaskService {

    List<Task> findAll();

    List<Task> findBySearchCriteria(Specification spec);

    List<Task> findTaskByLabel(String label);

    Task findById(Long taskId);

    Task createTask(Task task);

    Task updateTask(Task task, Long taskId);

    void deleteTask(Long taskId);

}
