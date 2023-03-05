package com.lcc.cloudservice.controller;

import com.lcc.cloudservice.common.R;
import com.lcc.cloudservice.entity.User;
import com.lcc.cloudservice.service.IUserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpSession;
import java.util.List;

/**
 * <p>
 *  前端控制器
 * </p>
 *
 * @author lcc
 * @since 2022-10-12
 */
@RestController
@RequestMapping("/user")
public class UserController {
    @Autowired
    IUserService iUserService;

    @GetMapping("/list")
    public R getUserList(){
        List<User> list = iUserService.getUserList();
        return R.ok().put("data", list);
    }

    @PostMapping("/add")
    public R setUser(@RequestBody User user){
        iUserService.addUser(user);
        return R.ok();
    }

    @GetMapping("/delete")
    public R setUser(@RequestBody int id){
        iUserService.deleteUserById(id);
        return R.ok();
    }

    @GetMapping("/login")
    public R setUser(@RequestBody User user, HttpServletRequest httpServletRequest){
        if(iUserService.verifyUser(user)){
            user = iUserService.getUserByName(user.getUsername());
            HttpSession session = httpServletRequest.getSession();
            session.setAttribute("user", user);
            return R.ok();
        }else{
            return R.error("username and password do not match");
        }

    }

    @GetMapping("/logout")
    public R setUser(HttpServletRequest httpServletRequest){
        HttpSession session = httpServletRequest.getSession();
        if(session.getAttribute("user") != null){
            session.removeAttribute("user");
        }
        return R.ok();
    }
}
