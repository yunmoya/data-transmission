package com.lcc.cloudservice.mapper;

import com.lcc.cloudservice.entity.VirtualMachine;
import com.baomidou.mybatisplus.core.mapper.BaseMapper;

/**
 * <p>
 *  Mapper 接口
 * </p>
 *
 * @author lcc
 * @since 2022-11-14
 */
public interface VirtualMachineMapper extends BaseMapper<VirtualMachine> {

    VirtualMachine getRestVMOne(int status, int configId);
}
