package com.example.javaspring.repository;

import org.springframework.data.jpa.repository.JpaRepository;

import com.example.javaspring.model.Task;


public interface TaskRepository extends JpaRepository<Task, Long> {
}
