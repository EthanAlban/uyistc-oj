package com.judgeserver.oj.entity;

import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@ApiModel(value = "Case对象",description = "测试案例对象")
public class TestCase {

    @ApiModelProperty(name = "TestCaseId" , value = "测试案例ID",example = "1")
    Long TestCaseId;

    @ApiModelProperty(name = "Input",value = "测试案例输入")
    String Input;

    @ApiModelProperty(name = "Output",value = "测试案例输出")
    String Output;

    public Long getTestCaseId() {
        return TestCaseId;
    }

    public void setTestCaseId(Long testCaseId) {
        TestCaseId = testCaseId;
    }

    public String getInput() {
        return Input;
    }

    public void setInput(String input) {
        Input = input;
    }

    public String getOutput() {
        return Output;
    }

    public void setOutput(String output) {
        Output = output;
    }
}
