package com.lcc.cloudservice.vos;

import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import lombok.Data;

@Data
@ApiModel(value = "用户请求对象")
public class ReqVo {
    @ApiModelProperty("用户id")
    String userId;
    @ApiModelProperty("请求配置（cpu核心数）")
    Integer config;
    @ApiModelProperty("分配时长（日）")
    Long duration;
}
