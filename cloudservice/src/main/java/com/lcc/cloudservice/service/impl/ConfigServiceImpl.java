package com.lcc.cloudservice.service.impl;

import com.baomidou.mybatisplus.core.conditions.query.QueryWrapper;
import com.lcc.cloudservice.entity.Config;
import com.lcc.cloudservice.mapper.ConfigMapper;
import com.lcc.cloudservice.service.IConfigService;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import org.springframework.stereotype.Service;

import javax.annotation.Resource;
import java.util.List;

/**
 * <p>
 *  服务实现类
 * </p>
 *
 * @author lcc
 * @since 2022-11-14
 */
@Service
public class ConfigServiceImpl extends ServiceImpl<ConfigMapper, Config> implements IConfigService {
    @Resource
    ConfigMapper configMapper;

    @Override
    public void add(Integer config) {
        if(checkExist(config)){
            return;
        }
        Config newConfig = new Config();
        newConfig.setCpuNum(config);
        configMapper.insert(newConfig);
    }

    @Override
    public boolean checkExist(Integer config) {
        List<Config> result = configMapper.selectList(new QueryWrapper<Config>().eq("cpu_num", config));
        return result.size() > 0;
    }

    @Override
    public void delete(Integer config) {
        Config result = configMapper.selectOne(new QueryWrapper<Config>().eq("cpu_num", config));
        if(result != null) {
            configMapper.deleteById(result.getId());
        }
    }

    @Override
    public float computePropotion() {
        int k = 0;
        List<Config> configList = configMapper.selectList(new QueryWrapper<Config>());
        for(Config config : configList) {
            k += config.getRestNum() * config.getCpuNum();
        }
        return k;
    }

    @Override
    public Config getConfig(Integer config) {
        return configMapper.selectOne(new QueryWrapper<Config>().eq("cpu_num", config));
    }

    @Override
    public void update(Config configResult) {
        configMapper.updateById(configResult);
    }
}
