package com.example.javaspring.repository;

import org.springframework.data.jpa.domain.Specification;
import java.util.ArrayList;
import java.util.List;

import com.example.javaspring.model.Task;

public class TaskSpecificationBuilder {

    private final List<SearchCriteria> params;

    public TaskSpecificationBuilder() {
        params = new ArrayList<SearchCriteria>();
    }

    public TaskSpecificationBuilder with(String key, String operation, Object value) {
        params.add(new SearchCriteria(key, operation, value));
        return this;
    }

    public Specification<Task> build() {
        if (params.size() == 0) {
            return null;
        }

        Specification<Task> result = new TaskSpecification(params.get(0));

        for (SearchCriteria param : params) {
            result = new TaskSpecification(param);
        }
        return result;
    }

}
