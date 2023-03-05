package com.lcc.cloudservice.entity;

import com.baomidou.mybatisplus.annotation.IdType;
import com.baomidou.mybatisplus.annotation.TableId;
import java.io.Serializable;
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
@ApiModel(value = "Config对象", description = "")
public class Config implements Serializable {

    private static final long serialVersionUID = 1L;

      @ApiModelProperty("配置id")
        @TableId(value = "id", type = IdType.AUTO)
      private Integer id;

      @ApiModelProperty("机器cpu核心数")
      private Integer cpuNum;

      @ApiModelProperty("机器剩余数量")
      private Integer restNum;

      @ApiModelProperty("机器总数")
      private Integer countNum;

    
    public Integer getId() {
        return id;
    }

      public void setId(Integer id) {
          this.id = id;
      }
    
    public Integer getCpuNum() {
        return cpuNum;
    }

      public void setCpuNum(Integer cpuNum) {
          this.cpuNum = cpuNum;
      }
    
    public Integer getRestNum() {
        return restNum;
    }

      public void setRestNum(Integer restNum) {
          this.restNum = restNum;
      }
    
    public Integer getCountNum() {
        return countNum;
    }

      public void setCountNum(Integer countNum) {
          this.countNum = countNum;
      }

    @Override
    public String toString() {
        return "Config{" +
              "id=" + id +
                  ", cpuNum=" + cpuNum +
                  ", restNum=" + restNum +
                  ", countNum=" + countNum +
              "}";
    }
}
