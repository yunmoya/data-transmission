package com.lcc.cloudservice.vos;


import io.swagger.annotations.ApiModelProperty;
import lombok.Data;

import java.time.LocalDateTime;

@Data
public class AssignVo {
    @ApiModelProperty("被分配的虚拟机id")
    private Long vmId;

    @ApiModelProperty("用户id")
    private String userId;

    @ApiModelProperty("被分配的虚拟机名")
    private String vmName;

    @ApiModelProperty("分配时间")
    private LocalDateTime assignTime;

    @ApiModelProperty("结束时间")
    private LocalDateTime endTime;
}
