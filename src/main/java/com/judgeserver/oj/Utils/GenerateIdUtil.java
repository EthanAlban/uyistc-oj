package com.judgeserver.oj.Utils;

/**
 * 雪花算法生成ID
 */
public class GenerateIdUtil {

    private final static int WORK_ID_BITS = 10;

    private final static int SEQUENCE_ID_BITS = 12;

    private static Long LAST_TIME_STAMP = -1L;

    private static Long WORK_ID = 12L;

    private static Long SEQUENCE_ID = 1L;

    public static synchronized Long getNextId(){
        Long TIME_STAMP = System.currentTimeMillis();
        if(TIME_STAMP == LAST_TIME_STAMP){
            if(SEQUENCE_ID >= 4096L){
                while(LAST_TIME_STAMP!=System.currentTimeMillis()){

                }
                SEQUENCE_ID = 1L;
                LAST_TIME_STAMP = System.currentTimeMillis();
            }
            LAST_TIME_STAMP = TIME_STAMP;
            SEQUENCE_ID++;
        }else{
            LAST_TIME_STAMP = TIME_STAMP;
            SEQUENCE_ID = 1L;
        }
        return LAST_TIME_STAMP<<(WORK_ID_BITS+SEQUENCE_ID_BITS)+WORK_ID<<SEQUENCE_ID_BITS+SEQUENCE_ID;
    }
}
