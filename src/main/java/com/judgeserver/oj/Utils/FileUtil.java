package com.judgeserver.oj.Utils;

import org.springframework.web.multipart.MultipartFile;

import java.io.File;
import java.io.IOException;

public class FileUtil {

    private final static String REL_Path = "\\src\\main\\resources\\";

    private final static String USER_DIR = "user.dir";
    public synchronized static String writeFile(MultipartFile file,String dictory){
        //String[] str = file.getOriginalFilename().split("\\.");
        String fileName = System.getProperty(USER_DIR)+REL_Path+dictory+"\\"+System.currentTimeMillis()+"_"+file.getOriginalFilename();
        File destFile = new File(fileName);
        try{
            if(destFile.exists()){
                destFile.createNewFile();
            }
            file.transferTo(destFile);
        }catch (IOException e){
            e.printStackTrace();
            return "error";
        }
        return fileName;
    }

}
