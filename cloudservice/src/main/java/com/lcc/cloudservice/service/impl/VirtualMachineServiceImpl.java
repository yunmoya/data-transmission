package com.lcc.cloudservice.service.impl;

import com.baomidou.mybatisplus.core.conditions.query.QueryWrapper;
import com.baomidou.mybatisplus.core.conditions.update.UpdateWrapper;
import com.lcc.cloudservice.entity.Cloud;
import com.lcc.cloudservice.entity.Config;
import com.lcc.cloudservice.entity.VirtualMachine;
import com.lcc.cloudservice.mapper.VirtualMachineMapper;
import com.lcc.cloudservice.service.IConfigService;
import com.lcc.cloudservice.service.IVirtualMachineService;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import com.lcc.cloudservice.uid.IdWorker;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import javax.annotation.Resource;
import java.time.LocalDateTime;

/**
 * <p>
 *  服务实现类
 * </p>
 *
 * @author lcc
 * @since 2022-11-14
 */
@Service
public class VirtualMachineServiceImpl extends ServiceImpl<VirtualMachineMapper, VirtualMachine> implements IVirtualMachineService {
    @Resource
    VirtualMachineMapper virtualMachineMapper;

    @Autowired
    IConfigService iConfigService;

    @Autowired
    IdWorker idWorker;

    @Value("${csp.name}")
    private String cspName;

    @Override
    @Transactional
    public void add(Integer num, Integer config) {
        Config configResult = iConfigService.getConfig(config);
        if(configResult != null) {
            int configId = configResult.getId();
            for (int i = 0; i < num; i++) {
                VirtualMachine vm = new VirtualMachine();
                vm.setStatus(0);
                vm.setCreateTime(LocalDateTime.now());
                vm.setId(idWorker.nextId());
                vm.setName(cspName.toLowerCase()+"_"+vm.getId());
                vm.setConfigId(configId);
                virtualMachineMapper.insert(vm);
            }
            configResult.setCountNum(configResult.getCountNum() + num);
            configResult.setRestNum(configResult.getRestNum() + num);
            iConfigService.update(configResult);
        }

    }

    @Override
    @Transactional
    public void delete(Long vmId) {
        VirtualMachine vm = virtualMachineMapper.selectById(vmId);
        if(vm != null && vm.getStatus() == 0){
            Config config = iConfigService.getById(vm.getConfigId());
            config.setCountNum(config.getCountNum() - 1);
            config.setRestNum(config.getRestNum() - 1);
            iConfigService.update(config);
            virtualMachineMapper.deleteById(vmId);
        }
    }

    @Override
    public VirtualMachine getRestVM(Integer configId) {
        return virtualMachineMapper.getRestVMOne(0, configId);
    }
}
