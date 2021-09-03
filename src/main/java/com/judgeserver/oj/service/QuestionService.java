package com.judgeserver.oj.service;

import com.judgeserver.oj.entity.Question;
import org.springframework.stereotype.Service;

import java.util.HashMap;
import java.util.List;
import java.util.Map;


public interface QuestionService {

    public void updateOne(Question question);

    public void addOne(Question question);

    public void deleteOneByPid(Long pid);

    public void queryOne(Long pid);

    public List<Map<String,Object>> getCaseMsgByPid(Long pid);
}
