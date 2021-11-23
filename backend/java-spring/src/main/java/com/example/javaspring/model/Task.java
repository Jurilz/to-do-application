package com.example.javaspring.model;

import com.example.javaspring.util.DateTimeUtil;

import javax.persistence.*;
import java.util.Objects;

@Entity
@Table(name = "tasks")
public class Task {
    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private Long id;
    private String label;
    private String  date;
    private boolean done;

    public Task(String label) {
        this.label = label;
        this.date = DateTimeUtil.getCurrentDate();
        this.done = false;
    }

    public Task() { }

    public void setId(Long id) {
        this.id = id;
    }

    public Long getId() {
        return id;
    }

    public String getLabel() {
        return label;
    }

    public void setLabel(String label) {
        this.label = label;
    }

    public String getDate() {
        return date;
    }

    public void setDate(String date) { this.date = date; };

    public Boolean getDone() {
        return done;
    }

    public void setDone(Boolean done) {
        this.done = done;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;

        if ( !(o instanceof Task)) return false;

        try {
            Task task = (Task) o;
            return Objects.equals(this.id, task.id) && Objects.equals(this.date, task.date) && Objects.equals(this.done, task.done);
        } catch (Exception e) {
            System.out.println(e.getMessage());
            return false;
        }
    }

    @Override
    public int hashCode() {
        return Objects.hash(this.id, this.label, this.date, this.done);
    }

    @Override
    public String toString() {
        return "Task {" + " id=" + this.id + " label=" + this.label + " date=" + this.date + " done=" + this.done + " }";
     }
}
