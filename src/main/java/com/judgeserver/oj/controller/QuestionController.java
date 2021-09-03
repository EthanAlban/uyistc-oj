package com.judgeserver.oj.controller;


import com.alibaba.fastjson.JSONObject;
import com.judgeserver.oj.Utils.GenerateIdUtil;
import com.judgeserver.oj.entity.Question;
import com.judgeserver.oj.service.QuestionService;
import io.swagger.annotations.*;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

@Api(tags = "问题管理")
@RestController
@RequestMapping("/question")
public class QuestionController {

    @Autowired
    QuestionService questionService;

    @ApiOperation(value = "获取测试案例",notes = "根据问题id获取到该问题下所有测试案例", httpMethod = "POST")
    @RequestMapping(value = "/getCaseMsgByPid")
    @ApiImplicitParams(
            @ApiImplicitParam(value = "问题ID",name = "pid",dataType = "long",paramType = "query")
    )
    public String getCaseMsgByPid(Long pid){
        List<Map<String,Object>> list = questionService.getCaseMsgByPid(pid);
        HashMap map = new HashMap();
        map.put("result",list);
        return JSONObject.toJSONString(map);
    }


    @ApiOperation(value = "删除问题",notes = "根据问题Id删除问题",httpMethod = "GET")
    @ApiImplicitParams(
            @ApiImplicitParam(value = "问题ID",name = "pid",dataType = "long",paramType = "query")
    )
    @RequestMapping("/deleteOne")
    public String deleteOne(Long pid){
        HashMap map = new HashMap();
        questionService.deleteOneByPid(pid);
        map.put("result","OK");
        return JSONObject.toJSONString(map);
    }

    @ApiOperation(value = "添加问题",httpMethod = "POST")
    @RequestMapping("/addOne")
    public String addOne(Question question){
        HashMap map = new HashMap();
        question.setPid(GenerateIdUtil.getNextId());
        questionService.addOne(question);
        map.put("result","OK");
        return JSONObject.toJSONString(map);
    }

    @ApiOperation(value = "修改问题",httpMethod = "POST")
    @RequestMapping("/updateOne")
    public String updateOne(Question question){
        questionService.updateOne(question);
        HashMap map = new HashMap();
        map.put("result","OK");
        return JSONObject.toJSONString(map);
    }
}
