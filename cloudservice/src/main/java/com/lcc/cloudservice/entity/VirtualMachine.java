package com.lcc.cloudservice.entity;

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
@TableName("virtual_machine")
@ApiModel(value = "VirtualMachine对象", description = "")
public class VirtualMachine implements Serializable {

    private static final long serialVersionUID = 1L;

      @ApiModelProperty("机器id")
        private Long id;

      @ApiModelProperty("0:未被分配 1:已分配")
      private Integer status;

      @ApiModelProperty("机器名")
      private String name;

      @ApiModelProperty("创建时间")
      private LocalDateTime createTime;

      @ApiModelProperty("配置id")
      private Integer configId;

      @ApiModelProperty("机器价格")
      private Integer cost;

    
    public Long getId() {
        return id;
    }

      public void setId(Long id) {
          this.id = id;
      }
    
    public Integer getStatus() {
        return status;
    }

      public void setStatus(Integer status) {
          this.status = status;
      }
    
    public String getName() {
        return name;
    }

      public void setName(String name) {
          this.name = name;
      }
    
    public LocalDateTime getCreateTime() {
        return createTime;
    }

      public void setCreateTime(LocalDateTime createTime) {
          this.createTime = createTime;
      }
    
    public Integer getConfigId() {
        return configId;
    }

      public void setConfigId(Integer configId) {
          this.configId = configId;
      }
    
    public Integer getCost() {
        return cost;
    }

      public void setCost(Integer cost) {
          this.cost = cost;
      }

    @Override
    public String toString() {
        return "VirtualMachine{" +
              "id=" + id +
                  ", status=" + status +
                  ", name=" + name +
                  ", createTime=" + createTime +
                  ", configId=" + configId +
                  ", cost=" + cost +
              "}";
    }
}
