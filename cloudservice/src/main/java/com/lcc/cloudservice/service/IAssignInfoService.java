package com.lcc.cloudservice.service;

import com.lcc.cloudservice.entity.AssignInfo;
import com.baomidou.mybatisplus.extension.service.IService;
import com.lcc.cloudservice.vos.AssignVo;

/**
 * <p>
 *  服务类
 * </p>
 *
 * @author lcc
 * @since 2022-11-14
 */
public interface IAssignInfoService extends IService<AssignInfo> {
    AssignVo addInfo(String userId, Integer config, Long duration);

    void revoke(String userId, Long vmId);
}
