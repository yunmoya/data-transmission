package com.lcc.cloudservice.controller;

import com.lcc.cloudservice.common.R;
import com.lcc.cloudservice.service.IConfigService;
import io.swagger.annotations.*;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;
import org.springframework.stereotype.Controller;

/**
 * <p>
 *  前端控制器
 * </p>
 *
 * @author lcc
 * @since 2022-11-14
 */
@RestController
@RequestMapping("/config")
public class ConfigController {

    @Autowired
    IConfigService iConfigService;

    @ApiOperation("添加一种服务器配置")
    @ApiResponses({
            @ApiResponse(code = 200, message = "添加成功"),
            @ApiResponse(code = 400, message = "请求参数错误"),
            @ApiResponse(code = 500, message = "系统错误")
    })
    @ApiImplicitParams({
            @ApiImplicitParam(paramType = "query", name = "config", dataType = "Integer", required = true, value = "请求配置（cpu核心数）")
    })
    @GetMapping("/add")
    public R addConfig(@RequestParam Integer config) {
        if(config == null || config < 0 || iConfigService.checkExist(config)) {
            return R.error(400, "非法参数");
        }
        iConfigService.add(config);
        return R.ok();
    }

    @ApiOperation("删除一种服务器配置")
    @ApiResponses({
            @ApiResponse(code = 200, message = "添加成功"),
            @ApiResponse(code = 400, message = "请求参数错误"),
            @ApiResponse(code = 500, message = "系统错误")
    })
    @ApiImplicitParams({
            @ApiImplicitParam(paramType = "query", name = "config", dataType = "Integer", required = true, value = "请求配置（cpu核心数）")
    })
    @GetMapping("/delete")
    public R deleteConfig(@RequestParam Integer config) {
        if(config == null || config < 0) {
            return R.error(400, "非法参数");
        }
        iConfigService.delete(config);
        return R.ok();
    }

    @ApiOperation("查询指定服务器配置是否存在")
    @ApiResponses({
            @ApiResponse(code = 200, message = "添加成功"),
            @ApiResponse(code = 400, message = "请求参数错误"),
            @ApiResponse(code = 500, message = "系统错误")
    })
    @ApiImplicitParams({
            @ApiImplicitParam(paramType = "path", name = "config", dataType = "Integer", required = true, value = "请求配置（cpu核心数）")
    })
    @GetMapping("/exist/{config}")
    public R checkExist(@PathVariable("config") Integer config) {
        if(config == null || config < 0) {
            return R.error(400, "非法参数");
        }
        if(iConfigService.checkExist(config)){
            return R.ok().put("Exist", true);
        } else{
            return R.ok().put("Exist", false);
        }
    }

    @ApiOperation("计算提供商proportion")
    @ApiResponses({
            @ApiResponse(code = 200, message = "计算成功"),
            @ApiResponse(code = 500, message = "系统错误")
    })
    @GetMapping("/proportion")
    public R getProportion() {
        float proportion = iConfigService.computePropotion();
        return R.ok().put("Proportion", proportion);
    }
}
