package com.judgeserver.oj.service.Impl;

import com.judgeserver.oj.Utils.FileUtil;
import com.judgeserver.oj.Utils.GenerateIdUtil;
import com.judgeserver.oj.Utils.StringUtil;
import com.judgeserver.oj.dao.TestCaseMapper;
import com.judgeserver.oj.entity.TestCase;
import com.judgeserver.oj.service.TestCaseService;
import org.springframework.stereotype.Service;
import org.springframework.web.multipart.MultipartFile;

import javax.annotation.Resource;
import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;

@Service
public class TestCaseServiceImpl implements TestCaseService {

    @Resource
    TestCaseMapper testCaseMapper;


    BufferedReader bufferedReader = null;

    private final static String CASE_FORMAT_ERROR = "案例输入输出格式错误";


    @Override
    public void addMultipleCase(List<TestCase> list) {
        testCaseMapper.addMultipleCase(list);
    }

    @Override
    public HashMap queryOne(Long caseId) {
        return testCaseMapper.queryOne(caseId);
    }

    @Override
    public void addCase(TestCase testCase) {
        testCaseMapper.addCase(testCase);
    }

    @Override
    public void addRelation(Long pid, Long caseId) {
        testCaseMapper.addRelation(pid,caseId);
    }

    @Override
    public void deleteCase(Long caseId) {
        testCaseMapper.deleteCase(caseId);
    }

    @Override
    public void deleteMultipleCase(Long pid) {
        testCaseMapper.deleteCase(pid);
    }

    @Override
    public void updateCase(TestCase testCase) {
        testCaseMapper.updateCase(testCase);
    }

    @Override
    public String analysisCaseText(MultipartFile file,Long pid){
        String fileName = FileUtil.writeFile(file,"temp");
        File tempFile = new File(fileName);
        try{
            bufferedReader = new BufferedReader(new FileReader(tempFile));
            String line = null;
            TestCase testCase = new TestCase();
            List<String> list =null;
            while ((line = bufferedReader.readLine()) != null) {
                list = StringUtil.toArrays(line);
                if(list.size() == 2){
                    Long caseId = GenerateIdUtil.getNextId();
                    testCase.setInput(list.get(0));
                    testCase.setOutput(list.get(1));
                    testCase.setTestCaseId(caseId);
                    testCaseMapper.addCase(testCase);
                    testCaseMapper.addRelation(pid,caseId);
                }else{
                    return CASE_FORMAT_ERROR;
                }
            }
            //tempFile.delete();
        }catch (IOException e){
            e.printStackTrace();
        }

        return "OK";
    }
}
