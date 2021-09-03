package com.judgeserver.oj.service.Impl;

import com.judgeserver.oj.dao.QuestionMapper;
import com.judgeserver.oj.entity.Question;
import com.judgeserver.oj.service.QuestionService;
import org.springframework.stereotype.Service;

import javax.annotation.Resource;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

@Service
public class QuestionServiceImpl implements QuestionService {

    @Resource
    QuestionMapper questionMapper;

    @Override
    public void addOne(Question question){
        questionMapper.addOne(question);
    }

    @Override
    public void deleteOneByPid(Long pid) {
        questionMapper.deleteOneByPid(pid);
    }

    @Override
    public void queryOne(Long pid) {
        questionMapper.queryOne(pid);
    }

    @Override
    public void updateOne(Question question) {
        questionMapper.updateOne(question);
    }

    @Override
    public List<Map<String,Object>> getCaseMsgByPid(Long pid) {
        return questionMapper.getCaseMsgByPid(pid);
    }
}
