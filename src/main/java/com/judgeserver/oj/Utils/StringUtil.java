package com.judgeserver.oj.Utils;

import java.util.ArrayList;
import java.util.List;

public class StringUtil {

    /**
     * 将字符串解析成[]形式字符串list
     * @param str
     * @return List<String>
     */
    public static List<String> toArrays(String str){
        int left=0,right=0,index=0;
        List<String> list = new ArrayList<>();
        for(int i = 0;i<str.length();i++){
            if(str.charAt(i) == '['){
                if(left == 0){
                    index = i;
                }
                left++;
            }else if(str.charAt(i) == ']'){
                right++;
            }
            if(left !=0 && left==right){
                list.add(str.substring(index,i+1));
                left=0;
                right=0;
            }
        }
        return list;
    }
}
