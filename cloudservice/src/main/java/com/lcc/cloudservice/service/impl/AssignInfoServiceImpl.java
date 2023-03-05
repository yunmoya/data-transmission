package com.lcc.cloudservice.service.impl;

import com.baomidou.mybatisplus.core.conditions.query.QueryWrapper;
import com.lcc.cloudservice.entity.AssignInfo;
import com.lcc.cloudservice.entity.Config;
import com.lcc.cloudservice.entity.VirtualMachine;
import com.lcc.cloudservice.mapper.AssignInfoMapper;
import com.lcc.cloudservice.service.IAssignInfoService;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import com.lcc.cloudservice.service.IConfigService;
import com.lcc.cloudservice.service.IVirtualMachineService;
import com.lcc.cloudservice.vos.AssignVo;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import javax.annotation.Resource;
import java.time.LocalDateTime;
import java.time.temporal.ChronoUnit;

/**
 * <p>
 *  服务实现类
 * </p>
 *
 * @author lcc
 * @since 2022-11-14
 */
@Service
public class AssignInfoServiceImpl extends ServiceImpl<AssignInfoMapper, AssignInfo> implements IAssignInfoService {
    @Resource
    AssignInfoMapper assignInfoMapper;

    @Autowired
    IConfigService iConfigService;

    @Autowired
    IVirtualMachineService iVirtualMachineService;

    @Override
    @Transactional
    public AssignVo addInfo(String userId, Integer config, Long duration) {
        AssignInfo assignInfo = new AssignInfo();
        assignInfo.setUserId(userId);
        Config configResult = iConfigService.getConfig(config);
        VirtualMachine vm = iVirtualMachineService.getRestVM(configResult.getId());
        assignInfo.setVmId(vm.getId());
        assignInfo.setAssignTime(LocalDateTime.now());
        assignInfo.setEndTime(assignInfo.getAssignTime().plus(duration, ChronoUnit.DAYS));
        assignInfo.setStatus(0);
        assignInfoMapper.insert(assignInfo);
        vm.setStatus(1);
        iVirtualMachineService.updateById(vm);
        configResult.setRestNum(configResult.getRestNum() - 1);
        iConfigService.updateById(configResult);
        AssignVo assignVo = new AssignVo();
        assignVo.setAssignTime(assignInfo.getAssignTime());
        assignVo.setEndTime(assignInfo.getEndTime());
        assignVo.setUserId(userId);
        assignVo.setVmId(vm.getId());
        assignVo.setVmName(vm.getName());
        return assignVo;
    }

    @Override
    @Transactional
    public void revoke(String userId, Long vmId) {
        AssignInfo assignInfo = assignInfoMapper.selectOne(new QueryWrapper<AssignInfo>().eq("user_id", userId).eq("vm_id", vmId));
        if(assignInfo.getStatus() == 0){
            VirtualMachine vm = iVirtualMachineService.getById(vmId);
            Config config = iConfigService.getById(vm.getConfigId());
            assignInfo.setStatus(1);
            assignInfoMapper.updateById(assignInfo);
            vm.setStatus(0);
            iVirtualMachineService.updateById(vm);
            config.setRestNum(config.getRestNum()+1);
            iConfigService.updateById(config);
        }

    }
}
