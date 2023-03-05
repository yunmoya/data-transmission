package com.lcc.cloudservice.controller;

import com.lcc.cloudservice.common.R;
import com.lcc.cloudservice.service.IConfigService;
import com.lcc.cloudservice.service.IVirtualMachineService;
import io.swagger.annotations.*;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
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
@RequestMapping("/virtualMachine")
public class VirtualMachineController {
    @Autowired
    IVirtualMachineService iVirtualMachineService;

    @Autowired
    IConfigService iConfigService;

    @Value("${csp.name}")
    private String cspName;

    @ApiOperation("批量添加服务器")
    @ApiResponses({
            @ApiResponse(code = 200, message = "添加成功"),
            @ApiResponse(code = 400, message = "参数错误"),
            @ApiResponse(code = 500, message = "系统错误")
    })
    @ApiImplicitParams({
            @ApiImplicitParam(paramType = "query", name = "config", dataType = "Integer", required = true, value = "请求配置（cpu核心数）"),
            @ApiImplicitParam(paramType = "query", name = "num", dataType = "Integer", required = true, value = "服务器数量")
    })
    @GetMapping("/add")
    public R add(@RequestParam("config") Integer config, @RequestParam Integer num){
        if(!iConfigService.checkExist(config)) {
            // 添加配置
            return R.error(400, "无效配置");
        }
        if(num != null)
            iVirtualMachineService.add(num, config);
        return R.ok();
    }

    @ApiOperation("根据ID删除服务器")
    @ApiResponses({
            @ApiResponse(code = 200, message = "删除成功")
    })
    @ApiImplicitParams({
            @ApiImplicitParam(paramType = "path", name = "vmId", dataType = "Integer", required = true, value = "服务器ID")
    })
    @GetMapping("/delete/{vmId}")
    public R delete(@PathVariable Long vmId){
        iVirtualMachineService.delete(vmId);
        return R.ok();
    }
}
