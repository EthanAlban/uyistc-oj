package com.judgeserver.oj.dao;

import com.judgeserver.oj.entity.Question;
import org.apache.ibatis.annotations.Mapper;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

@Mapper
public interface QuestionMapper {

    public void updateOne(Question question);

    public void addOne(Question question);

    public void deleteOneByPid(Long pid);

    public void queryOne(Long pid);

    public List<Map<String,Object>> getCaseMsgByPid(Long pid);
}
