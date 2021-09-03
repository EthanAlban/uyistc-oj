package com.judgeserver.oj.dao;

import com.judgeserver.oj.entity.TestCase;
import org.apache.ibatis.annotations.Mapper;

import java.util.HashMap;
import java.util.List;
@Mapper
public interface TestCaseMapper {


    public void addMultipleCase(List<TestCase> list);

    public void addCase(TestCase testCase);

    public void addRelation(Long pid,Long caseId);

    public void deleteCase(Long caseId);

    public void updateCase(TestCase testCase);

    public HashMap queryOne(Long caseId);

    public void deleteMultipleCase(Long pid);
}
