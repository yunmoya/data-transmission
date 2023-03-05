package com.lcc.cloudservice;

import com.baomidou.mybatisplus.generator.FastAutoGenerator;
import org.mybatis.spring.annotation.MapperScan;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
@MapperScan("com.lcc.cloudservice.mapper")
public class CloudServiceApplication {

	public static void main(String[] args) {
		SpringApplication.run(CloudServiceApplication.class, args);
//		FastAutoGenerator.create("jdbc:mysql://10.1.0.133:3306/cloudservice?serverTimezone=Asia/Shanghai", "root", "bxsmm133")
//				.globalConfig(builder -> {
//					builder.author("lcc") // 设置作者
//							.enableSwagger(); // 开启 swagger 模式
//				})
//				.packageConfig(builder -> {
//					builder.parent("com.lcc"); // 设置父包名
//				})
//				.execute();

	}

}
