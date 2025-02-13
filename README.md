โปรเจ็กต์นี้เป็นระบบ **Kanban Board** ที่ช่วยให้ผู้ใช้สามารถจัดการและติดตามงานต่างๆ โดยสามารถโยกย้าย **Ticket** หรือ **Card** ได้อย่างอิสระ ซึ่งถูกพัฒนาขึ้นมาโดยใช้เทคโนโลยี **Frontend** ที่เชื่อมต่อกับ **Backend** API

## Frontend
- [Drive Frontend](https://drive.google.com/file/d/1Sl6x-xQim8ODSaEiMTiFEV-gx7vnG6oH/view?usp=sharing)

## API Endpoints

โปรเจ็กต์นี้ใช้ API ที่สามารถเข้าถึงได้ผ่าน `/api/v1` โดยมีรายละเอียดดังนี้:

### 1. Create Ticket
- **POST** `/ticket`
- สร้าง **Ticket** ใหม่

### 2. Get All Tickets
- **GET** `/ticket`
- ดึงข้อมูลทั้งหมดของ **Tickets**

### 3. Get Ticket By ID
- **GET** `/ticket/:id`
- ดึงข้อมูลของ **Ticket** ตาม ID ที่กำหนด

### 4. Update Ticket By ID
- **PUT** `/ticket/:id`
- อัพเดทข้อมูลของ **Ticket** ที่มี ID ที่กำหนด

## Features
- **Kanban Board** ที่ช่วยในการจัดการงานแบบยืดหยุ่น
- สามารถลากและย้าย **Ticket** ไปยังสถานะต่างๆ ได้ตามต้องการ
- ใช้งานง่าย และรองรับการจัดการหลายๆ งานพร้อมกัน

## เทคโนโลยีที่ใช้
- **Frontend**: React.js
- **Backend**: Golang (Gin Framework)
- **Database**: PostgreSQL (ฐานข้อมูลที่ใช้)

## ติดต่อ
- **ชื่อ**: นายจักรพัชร คัชชาพงษ์
- **อีเมล**: katchapong_j@silpakorn.edu
- **มหาวิทยาลัย**: มหาวิทยาลัยศิลปากร
