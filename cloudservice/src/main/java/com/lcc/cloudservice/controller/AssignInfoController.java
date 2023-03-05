package com.lcc.cloudservice.controller;

import com.lcc.cloudservice.common.R;
import com.lcc.cloudservice.entity.AssignInfo;
import com.lcc.cloudservice.entity.Cloud;
import com.lcc.cloudservice.entity.VirtualMachine;
import com.lcc.cloudservice.service.IAssignInfoService;
import com.lcc.cloudservice.service.IConfigService;
import com.lcc.cloudservice.service.IVirtualMachineService;
import com.lcc.cloudservice.vos.AssignVo;
import com.lcc.cloudservice.vos.ReqVo;
import io.swagger.annotations.*;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.util.ObjectUtils;
import org.springframework.web.bind.annotation.*;
import org.springframework.stereotype.Controller;

import javax.swing.*;

/**
 * <p>
 *  前端控制器
 * </p>
 *
 * @author lcc
 * @since 2022-11-14
 */
@RestController
@RequestMapping("/assignInfo")
public class AssignInfoController {
    @Autowired
    IConfigService iConfigService;

    @Autowired
    IAssignInfoService iAssignInfoService;

    @Value("${csp.name}")
    private String cspName;

    @ApiOperation("分配服务器")
    @ApiResponses({
            @ApiResponse(code = 200, message = "分配成功"),
            @ApiResponse(code = 400, message = "请求参数错误"),
            @ApiResponse(code = 500, message = "系统错误")
    })
    @PostMapping("/assign")
    public R getRestConfig(@RequestBody ReqVo reqVo) {
        if(ObjectUtils.isEmpty(reqVo)) {
            return R.error(400, "请求为空");
        }
        if(reqVo.getConfig() != null &&!iConfigService.checkExist(reqVo.getConfig())) {
            return R.error(500, "分配失败，服务商" + cspName + "资源不足。");
        }
        AssignVo assignVo = iAssignInfoService.addInfo(reqVo.getUserId(), reqVo.getConfig(), reqVo.getDuration());
        return R.ok().put("data", assignVo);
    }

    @ApiOperation("回收服务器")
    @ApiResponses({
            @ApiResponse(code = 200, message = "回收成功"),
            @ApiResponse(code = 500, message = "系统错误")
    })
    @ApiImplicitParams({
            @ApiImplicitParam(paramType = "path", name = "userId", dataType = "String", required = true, value = "用户ID"),
            @ApiImplicitParam(paramType = "query", name = "vmId", dataType = "Long", required = true, value = "被撤销服务器id")
    })

    @GetMapping("/revoke/{userId}")
    public R assign(@PathVariable String userId, @RequestParam Long vmId){
        iAssignInfoService.revoke(userId, vmId);
        return R.ok();
    }
}
