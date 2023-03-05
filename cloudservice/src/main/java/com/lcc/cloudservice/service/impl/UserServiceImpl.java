package com.lcc.cloudservice.service.impl;

import com.baomidou.mybatisplus.core.conditions.query.QueryWrapper;
import com.lcc.cloudservice.entity.User;
import com.lcc.cloudservice.mapper.UserMapper;
import com.lcc.cloudservice.service.IUserService;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

/**
 * <p>
 *  服务实现类
 * </p>
 *
 * @author lcc
 * @since 2022-10-12
 */
@Service
public class UserServiceImpl extends ServiceImpl<UserMapper, User> implements IUserService {
    @Autowired
    UserMapper userMapper;

    @Override
    public List<User> getUserList() {
        return userMapper.selectList(new QueryWrapper<User>().select("id", "username"));
    }

    @Override
    @Transactional
    public void addUser(User user) {
        userMapper.insert(user);
    }

    @Override
    @Transactional
    public void deleteUserById(int id) {
        userMapper.deleteById(id);
    }

    @Override
    public boolean verifyUser(User user) {
        User dbUser = userMapper.selectOne(new QueryWrapper<User>().eq("username", user.getUsername()));
        return dbUser.getPassword().equals(user.getPassword());
    }

    @Override
    public User getUserByName(String username) {
        return userMapper.selectOne(new QueryWrapper<User>().eq("username", username));
    }


}
