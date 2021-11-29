package com.example.javaspring.repository;

import com.example.javaspring.model.Task;
import org.springframework.data.jpa.domain.Specification;

import javax.persistence.criteria.CriteriaBuilder;
import javax.persistence.criteria.CriteriaQuery;
import javax.persistence.criteria.Predicate;
import javax.persistence.criteria.Root;
import java.util.function.Function;

public class TaskSpecification implements Specification<Task> {

    private final SearchCriteria criteria;

    final String DONE = "done";

    public TaskSpecification(final SearchCriteria criteria) {
        super();
        this.criteria = criteria;
    }

    public SearchCriteria getCriteria() {
        return this.criteria;
    }

    @Override
    public Predicate toPredicate(Root<Task> root, CriteriaQuery<?> query, CriteriaBuilder builder) {
       if (criteria.getOperation().equalsIgnoreCase(">")) {
           return builder.greaterThanOrEqualTo(
                   root.get(criteria.getKey()), criteria.getValue().toString());
       }

       else if (criteria.getOperation().equalsIgnoreCase("<")) {
           return builder.lessThanOrEqualTo(
                   root.get(criteria.getKey()), criteria.getValue().toString());
       }

       else if (criteria.getOperation().equalsIgnoreCase(":")) {
           if (root.get(criteria.getKey()).getJavaType() == String.class) {
               return builder.like(
                       root.get(criteria.getKey()), "%" + criteria.getValue() + "%");
           } else if (criteria.getKey().equalsIgnoreCase(DONE)) {
               if (criteria.getValue().toString().equalsIgnoreCase("false")) {
                   return builder.equal(root.get(criteria.getKey()), Boolean.FALSE);
               } else {
                   return builder.equal(root.get(criteria.getKey()), Boolean.TRUE);
               }
           } else {
               return builder.equal(root.get(criteria.getKey()), criteria.getValue());
           }
       }
       return null;
    }
}
