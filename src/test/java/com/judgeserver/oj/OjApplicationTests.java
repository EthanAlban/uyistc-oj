package com.judgeserver.oj;

import org.junit.jupiter.api.Test;
import org.springframework.boot.test.context.SpringBootTest;

@SpringBootTest
class OjApplicationTests {

    @Test
    void contextLoads() {

    }
    @Test
    public void test(){
        Long time = System.currentTimeMillis();
        System.out.println(time<<22);
    }

}
