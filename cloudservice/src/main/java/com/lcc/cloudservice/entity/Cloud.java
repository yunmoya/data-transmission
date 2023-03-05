package com.lcc.cloudservice.entity;

import com.baomidou.mybatisplus.annotation.IdType;
import com.baomidou.mybatisplus.annotation.TableId;
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
 * @since 2022-10-12
 */
@ApiModel(value = "Cloud对象", description = "")
public class Cloud implements Serializable {

    private static final long serialVersionUID = 1L;

      @TableId(value = "id", type = IdType.AUTO)
      private Long id;

      @ApiModelProperty("0:未被分配 1:已分配")
      private Integer status;

    private String name;

    private Integer userid;

    private LocalDateTime assignTime;

    private LocalDateTime createTime;

    
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
    
    public Integer getUserid() {
        return userid;
    }

      public void setUserid(Integer userid) {
          this.userid = userid;
      }
    
    public LocalDateTime getAssignTime() {
        return assignTime;
    }

      public void setAssignTime(LocalDateTime assignTime) {
          this.assignTime = assignTime;
      }
    
    public LocalDateTime getCreateTime() {
        return createTime;
    }

      public void setCreateTime(LocalDateTime createTime) {
          this.createTime = createTime;
      }

    @Override
    public String toString() {
        return "Cloud{" +
              "id=" + id +
                  ", status=" + status +
                  ", name=" + name +
                  ", userid=" + userid +
                  ", assignTime=" + assignTime +
                  ", createTime=" + createTime +
              "}";
    }
}
