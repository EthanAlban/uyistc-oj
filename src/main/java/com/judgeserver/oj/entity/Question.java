package com.judgeserver.oj.entity;

import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

@ApiModel(value = "Question对象",description = "问题对象")
public class Question {

    @ApiModelProperty(name = "pid",value = "问题id",example = "1")
    Long pid;

    @ApiModelProperty(name = "type",value = "问题输入输出类型",example = "1")
    String type;

    public Long getPid() {
        return pid;
    }

    public void setPid(Long pid) {
        this.pid = pid;
    }

    public String getType() {
        return type;
    }

    public void setType(String type) {
        this.type = type;
    }
}
