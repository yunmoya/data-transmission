package com.lcc.cloudservice.service;

import com.lcc.cloudservice.entity.User;
import com.baomidou.mybatisplus.extension.service.IService;

import java.util.List;

/**
 * <p>
 *  服务类
 * </p>
 *
 * @author lcc
 * @since 2022-10-12
 */
public interface IUserService extends IService<User> {

    List<User> getUserList();

    void addUser(User user);

    void deleteUserById(int id);

    boolean verifyUser(User user);

    User getUserByName(String username);
}
