package com.example.app;

import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.*;

/**
 * App类的单元测试
 */
public class AppTest {
    
    @Test
    public void testAdd() {
        assertEquals(4, App.add(2, 2), "2 + 2 应该等于 4");
    }
    
    @Test
    public void testAddNegative() {
        assertEquals(0, App.add(2, -2), "2 + (-2) 应该等于 0");
    }
} 