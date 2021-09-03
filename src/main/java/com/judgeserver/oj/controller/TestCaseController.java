package com.judgeserver.oj.controller;

import com.alibaba.fastjson.JSONObject;
import com.judgeserver.oj.Utils.GenerateIdUtil;
import com.judgeserver.oj.entity.TestCase;
import com.judgeserver.oj.service.TestCaseService;
import io.swagger.annotations.Api;
import io.swagger.annotations.ApiImplicitParam;
import io.swagger.annotations.ApiImplicitParams;
import io.swagger.annotations.ApiOperation;
import org.apache.ibatis.annotations.Param;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.multipart.MultipartFile;

import java.io.File;
import java.io.IOException;
import java.util.HashMap;
import java.util.List;

@Api(tags = "测试案例管理")
@RestController
@RequestMapping("/TestCase")
public class TestCaseController {
    @Autowired
    TestCaseService testCaseService;

    /**
     * 通过解析测试案例文件将测试案例存储到数据库中
     * @param file
     * @param pid
     * @return JSON字符串
     * @throws IOException
     */
    @ApiOperation(value = "添加测试案例",notes = "上传文件根据文件的解析来批量添加测试案例",httpMethod = "POST")
    @ApiImplicitParams({
            @ApiImplicitParam(name = "file",value = "测试案例文件，每一行中包括一个输入以及一个输出"),
            @ApiImplicitParam(name = "pid",value = "测试案例对应的问题id")
    })
    @RequestMapping("addMultipleCase")
    public String addMultipleCase(@Param("file") MultipartFile file,@Param("pid") Long pid) throws IOException {
        String res = testCaseService.analysisCaseText(file,pid);
        HashMap map = new HashMap();
        map.put("result",res);
        return JSONObject.toJSONString(map);
    }

    @ApiOperation(value = "添加一个测试案例",httpMethod = "POST")
    @ApiImplicitParams({
            @ApiImplicitParam(name = "pid",value = "测试案例对应的问题id")
    })
    @RequestMapping("addCase")
    public String addCase(TestCase testCase,@Param("pid") Long pid){
        testCase.setTestCaseId(GenerateIdUtil.getNextId());
        testCaseService.addCase(testCase);
        testCaseService.addRelation(pid,testCase.getTestCaseId());
        HashMap map = new HashMap();
        map.put("result","OK");
        return JSONObject.toJSONString(map);
    }

    @ApiOperation(value = "删除测试案例",httpMethod = "POST")
    @ApiImplicitParams({
            @ApiImplicitParam(name = "caseId",value = "测试案例ID")
    })
    @RequestMapping("deleteCase")
    public String deleteCase(@Param("caseId") Long caseId){
        testCaseService.deleteCase(caseId);
        HashMap map = new HashMap();
        map.put("result","OK");
        return JSONObject.toJSONString(map);
    }

    @ApiOperation(value = "修改测试案例",httpMethod = "POST")
    @RequestMapping("updateCase")
    public String updateCase(TestCase testCase){
        HashMap map = new HashMap();
        testCaseService.updateCase(testCase);
        map.put("result","OK");
        return JSONObject.toJSONString(map);
    }

    @ApiOperation(value = "查询测试案例",httpMethod = "POST")
    @ApiImplicitParams({
            @ApiImplicitParam(name = "caseId",value = "测试案例ID")
    })
    @RequestMapping("queryCase")
    public String queryCase(Long caseId){
        HashMap map = new HashMap();
        HashMap res = testCaseService.queryOne(caseId);
        map.put("result",res);
        return JSONObject.toJSONString(map);
    }
}
