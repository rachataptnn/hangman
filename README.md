# Hangman

เขียนโปรแกรมเพื่อเล่นเกม Hangman ใน Terminal/Console (Command Line Interface) โดยมีความสามารถ อย่างน้อยต่อไปนี้

เลือกหมวดหมู่ของคำได้ โดยเก็บคำในแต่ละหมวดหมู่แยกเป็นไฟล์ หมวดหมู่ละ 1 ไฟล์ โปรแกรมที่ส่งมาให้ตรวจจะต้องมีหมวดหมู่อย่างน้อย 2 หมวดหมู่ และมีคำในหมวดหมู่อย่างน้อยหมวดหมู่ละ 5 คำ

“คำ” ที่ใช้เล่นเกม อาจประกอบด้วยช่องว่าง ตัวเลข และสัญลักษณ์ แต่จะไม่เป็นส่วนหนึ่งของการเล่นเกม จะต้องแสดงผลทันที (เช่น คำอาจจะเป็น “Mong 31!!” ก็จะแสดงผลว่า \_ \_ \_ \_ 31!! เป็นต้น) การเล่นเกมจะใช้ตัวอักษรภาษาอังกฤษเท่านั้น โดยเป็น case-insensitive ทั้งหมด

คำแต่ละคำ จะต้องประกอบด้วย “คำใบ้”

มีการแสดงตัวอักษรที่เดาผิด

มีการคำนวนคะแนนจากการเล่นเดาคำในแต่ละรอบ โดยจะคำนวนคะแนนอย่างไร ขึ้นกับผู้ทำโจทย์ (คิดเกณฑ์/อัลกอริทึมการให้คะแนนเองได้เลย)

ในกรณีที่เดาซ้ำกับตัวอักษรที่เดาไปแล้ว ไม่ว่าจะถูกหรือผิด จะมีการแจ้งว่า already guessed และไม่มีการนับการเดา หรือคิดคะแนน

## ตัวอย่างการรันโปรแกรม

```
Select Category:
1. English Premiere League Teams (2018/2019)
2. Historical World Leaders
3. Famous Songs

> 2

Hint: “World War II”

_ _ _ _ _ _ _ _    score   0, remaining incorrect guess 10
> d
_ _ _ _ _ _ _ _    score   0, remaining incorrect guess  9, incorrect guessed: d
> r
_ _ r _ _ _ _ _    score  15, remaining incorrect guess  9, incorrect guessed: d
> d
already guessed
_ _ r _ _ _ _ _    score  15, remaining incorrect guess  9, incorrect guessed: d
>
```


## วิธีการทำโจทย์

1. เลือกภาษาที่ใช้เขียนโค้ด
2. พัฒนาโปรแกรม Hangman ตามโจทย์ที่ได้รับ
3. แนบคำอธิบาย หรือวิธีการสำหรับการใช้งานโปรแกรม

*** มีเวลาทำ 48 ชั่วโมง เริ่มนับเวลาจากที่ได้รับโจทย์ผ่านทาง Email

## Domains of Evaluation
Included but not limited to    
- Coding Style
- Programming Language Proficiency
- Object-Oriented / Functional Design and Programming