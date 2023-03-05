package com.lcc.cloudservice.service;

import com.lcc.cloudservice.entity.VirtualMachine;
import com.baomidou.mybatisplus.extension.service.IService;

/**
 * <p>
 *  服务类
 * </p>
 *
 * @author lcc
 * @since 2022-11-14
 */
public interface IVirtualMachineService extends IService<VirtualMachine> {

    void add(Integer num, Integer config);

    void delete(Long vmId);

    VirtualMachine getRestVM(Integer configId);
}
