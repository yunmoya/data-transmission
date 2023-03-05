package com.lcc.cloudservice.entity;

import com.baomidou.mybatisplus.annotation.IdType;
import com.baomidou.mybatisplus.annotation.TableId;
import com.baomidou.mybatisplus.annotation.TableName;
import java.io.Serializable;
import java.time.LocalDateTime;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;

/**
 * <p>
 * 
 * </p>
 *
 * @author lcc
 * @since 2022-11-14
 */
@TableName("assign_info")
@ApiModel(value = "AssignInfo对象", description = "")
public class AssignInfo implements Serializable {

    private static final long serialVersionUID = 1L;

      @ApiModelProperty("分配信息id")
        @TableId(value = "id", type = IdType.AUTO)
      private Long id;

      @ApiModelProperty("被分配的虚拟机id")
      private Long vmId;

      @ApiModelProperty("分配时间")
      private LocalDateTime assignTime;

      @ApiModelProperty("结束时间")
      private LocalDateTime endTime;

      @ApiModelProperty("分配状态 0：该分配依旧有效 1：该分配已过期")
      private Integer status;

      @ApiModelProperty("用户id")
      private String userId;

    
    public Long getId() {
        return id;
    }

      public void setId(Long id) {
          this.id = id;
      }
    
    public Long getVmId() {
        return vmId;
    }

      public void setVmId(Long vmId) {
          this.vmId = vmId;
      }
    
    public LocalDateTime getAssignTime() {
        return assignTime;
    }

      public void setAssignTime(LocalDateTime assignTime) {
          this.assignTime = assignTime;
      }
    
    public LocalDateTime getEndTime() {
        return endTime;
    }

      public void setEndTime(LocalDateTime endTime) {
          this.endTime = endTime;
      }
    
    public Integer getStatus() {
        return status;
    }

      public void setStatus(Integer status) {
          this.status = status;
      }
    
    public String getUserId() {
        return userId;
    }

      public void setUserId(String userId) {
          this.userId = userId;
      }

    @Override
    public String toString() {
        return "AssignInfo{" +
              "id=" + id +
                  ", vmId=" + vmId +
                  ", assignTime=" + assignTime +
                  ", endTime=" + endTime +
                  ", status=" + status +
                  ", userId=" + userId +
              "}";
    }
}
