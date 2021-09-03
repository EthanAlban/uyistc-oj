package com.judgeserver.oj.service;

import com.judgeserver.oj.entity.TestCase;
import org.springframework.stereotype.Service;
import org.springframework.web.multipart.MultipartFile;

import java.util.HashMap;
import java.util.List;


public interface TestCaseService {

    public void addMultipleCase(List<TestCase> list);

    public void addCase(TestCase testCase);

    public void addRelation(Long pid,Long caseId);

    public void deleteCase(Long caseId);

    public void updateCase(TestCase testCase);

    public HashMap queryOne(Long caseId);

    public void deleteMultipleCase(Long pid);

    public String analysisCaseText(MultipartFile file,Long pid);
}
