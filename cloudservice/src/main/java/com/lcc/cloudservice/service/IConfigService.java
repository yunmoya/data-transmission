package com.lcc.cloudservice.service;

import com.lcc.cloudservice.entity.Config;
import com.baomidou.mybatisplus.extension.service.IService;

/**
 * <p>
 *  服务类
 * </p>
 *
 * @author lcc
 * @since 2022-11-14
 */
public interface IConfigService extends IService<Config> {

    void add(Integer config);

    boolean checkExist(Integer config);

    void delete(Integer config);

    float computePropotion();

    Config getConfig(Integer config);

    void update(Config configResult);
}
