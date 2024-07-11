{
	"info": {
		"_postman_id": "12b273fe-7e8a-44c1-a5e3-af8e8d6501fd",
		"name": "KosLess API",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
		"_exporter_id": "36371785"
	},
	"item": [
		{
			"name": "Kos",
			"item": [
				{
					"name": "Create Kos",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"name\": \"Kosan Taman Hijau\",\r\n    \"address\": \"Jl. Kenari No. 10, Jakarta\",\r\n    \"roomCount\": 6,\r\n    \"coordinate\": \"-6.2088° S, 106.8456° E\",\r\n    \"description\": \"Kosan yang nyaman dan terjangkau, terletak di lingkungan yang tenang. Cocok untuk pelajar dan profesional muda.\",\r\n    \"rules\": \"Tidak boleh berisik setelah jam 10 malam, jaga kebersihan, dan hargai penghuni lain.\"\r\n}\r\n",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://localhost:4000/api/v1/kos",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "4000",
							"path": [
								"api",
								"v1",
								"kos"
							]
						}
					},
					"response": []
				},
				{
					"name": "Get Detail Kos By ID",
					"request": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "http://localhost:4000/api/v1/kos/0112f220-2bae-4a1b-8a40-672394e7616b",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "4000",
							"path": [
								"api",
								"v1",
								"kos",
								"0112f220-2bae-4a1b-8a40-672394e7616b"
							]
						}
					},
					"response": []
				},
				{
					"name": "Update Kos",
					"request": {
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"name\": \"Kosan Bunga Mawar\",\r\n    \"address\": \"Jl. Melati No. 5, Bandung\",\r\n    \"roomCount\": 6,\r\n    \"coordinate\": \"-6.9147° S, 107.6098° E\",\r\n    \"description\": \"Kosan yang nyaman dan bersih, terletak di lingkungan yang strategis. Cocok untuk mahasiswa dan karyawan.\",\r\n    \"rules\": \"Tidak boleh berisik setelah jam 10 malam, jaga kebersihan, dan hargai penghuni lain.\"\r\n}\r\n",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://localhost:4000/api/v1/kos/0112f220-2bae-4a1b-8a40-672394e7616b",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "4000",
							"path": [
								"api",
								"v1",
								"kos",
								"0112f220-2bae-4a1b-8a40-672394e7616b"
							]
						}
					},
					"response": []
				},
				{
					"name": "Delete Kos By ID",
					"request": {
						"method": "DELETE",
						"header": [],
						"url": {
							"raw": "http://localhost:4000/api/v1/kos/58de605a-9951-4adf-8ebb-64ce34513627",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "4000",
							"path": [
								"api",
								"v1",
								"kos",
								"58de605a-9951-4adf-8ebb-64ce34513627"
							]
						}
					},
					"response": []
				}
			]
		},
		{
			"name": "Room",
			"item": [
				{
					"name": "Create Room",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"kos_id\": \"d548feec-5ed5-4862-ba27-393117cb455e\",\r\n    \"name\": \"Kosan Taman Hijau\",\r\n    \"type\": \"VIP\",\r\n    \"description\": \"Kamar berkelas dengan fasilitas lengkap\",\r\n    \"avail\": \"open\",\r\n    \"price\": 2000000\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://localhost:4000/api/v1/room",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "4000",
							"path": [
								"api",
								"v1",
								"room"
							]
						}
					},
					"response": []
				},
				{
					"name": "Update Room",
					"request": {
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"kos_id\": \"0112f220-2bae-4a1b-8a40-672394e7616b\",\r\n    \"name\": \"Kos Bunga Mawar Merah\",\r\n    \"type\": \"VIP\",\r\n    \"description\": \"Kamar berkelas dengan fasilitas lengkap\",\r\n    \"avail\": \"open\",\r\n    \"price\": 2000000\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://localhost:4000/api/v1/room/fe625da3-3724-493e-a66e-a54db746f11c",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "4000",
							"path": [
								"api",
								"v1",
								"room",
								"fe625da3-3724-493e-a66e-a54db746f11c"
							]
						}
					},
					"response": []
				},
				{
					"name": "Get Room By ID",
					"request": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "http://localhost:4000/api/v1/room/73ca3fe7-e069-4b3a-b915-1a08c532c51d",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "4000",
							"path": [
								"api",
								"v1",
								"room",
								"73ca3fe7-e069-4b3a-b915-1a08c532c51d"
							]
						}
					},
					"response": []
				},
				{
					"name": "Get Room Price Lower Than",
					"request": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "http://localhost:4000/api/v1/room/budget?budget=2000000",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "4000",
							"path": [
								"api",
								"v1",
								"room",
								"budget"
							],
							"query": [
								{
									"key": "budget",
									"value": "2000000"
								}
							]
						}
					},
					"response": []
				},
				{
					"name": "Get Room By Avail",
					"request": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "http://localhost:4000/api/v1/room/availability/open",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "4000",
							"path": [
								"api",
								"v1",
								"room",
								"availability",
								"open"
							]
						}
					},
					"response": []
				},
				{
					"name": "Delete Room",
					"request": {
						"method": "DELETE",
						"header": [],
						"url": {
							"raw": "http://localhost:4000/api/v1/room/d5f92aa7-acc1-4e8e-af66-9cdd62d729bd",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "4000",
							"path": [
								"api",
								"v1",
								"room",
								"d5f92aa7-acc1-4e8e-af66-9cdd62d729bd"
							]
						}
					},
					"response": []
				},
				{
					"name": "Get All Rooms",
					"request": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "http://localhost:4000/api/v1/room/rooms",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "4000",
							"path": [
								"api",
								"v1",
								"room",
								"rooms"
							]
						}
					},
					"response": []
				}
			]
		},
		{
			"name": "User",
			"item": [
				{
					"name": "Register User",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"fullName\": \"Nicky Rahma Wati\",\r\n    \"username\": \"Nicky\",\r\n    \"password\": \"123\",\r\n    \"email\": \"nicky@gmail.com\",\r\n    \"phoneNumber\": \"987654321\",\r\n    \"photoProfile\": \"jpg\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://localhost:4000/api/v1/users/register",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "4000",
							"path": [
								"api",
								"v1",
								"users",
								"register"
							]
						}
					},
					"response": []
				},
				{
					"name": "Login User",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"username\": \"Desi\",\r\n    \"password\": \"123\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://localhost:4000/api/v1/users/login",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "4000",
							"path": [
								"api",
								"v1",
								"users",
								"login"
							]
						}
					},
					"response": []
				},
				{
					"name": "Get User By Id",
					"protocolProfileBehavior": {
						"disableBodyPruning": true
					},
					"request": {
						"auth": {
							"type": "bearer",
							"bearer": [
								{
									"key": "token",
									"value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJEZXNpIiwiZXhwIjoxNzIxMDY4NTIzLCJpYXQiOjE3MjA3MDg1MjMsImlkIjoiMjM5MWE5OWQtYTgwNS00MjQ3LTgwOWUtZmZjYzY1YWZmN2FkIiwidXNlcm5hbWUiOiJOaWNreSJ9.Tcx-rL1NOzbd23J1LHFjCKl549btiWCOcdTB8xO5hME",
									"type": "string"
								}
							]
						},
						"method": "GET",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://localhost:4000/api/v1/users/profile/2391a99d-a805-4247-809e-ffcc65aff7ad",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "4000",
							"path": [
								"api",
								"v1",
								"users",
								"profile",
								"2391a99d-a805-4247-809e-ffcc65aff7ad"
							]
						}
					},
					"response": []
				},
				{
					"name": "Update User",
					"request": {
						"auth": {
							"type": "bearer",
							"bearer": [
								{
									"key": "token",
									"value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJEZXNpIiwiZXhwIjoxNzIxMDY4MjE2LCJpYXQiOjE3MjA3MDgyMTYsImlkIjoiNTBhOTU5NDMtZGMxMi00YzFiLTg2ZDYtYTg1ZWQwNjRjNzI0IiwidXNlcm5hbWUiOiJEZXNpIn0.S_NYlIQhKtRdPmd-PLM4oUfTgKVnqYxoJ3X3ILA_R5Q",
									"type": "string"
								}
							]
						},
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"fullName\": \"Nicky RahmaWati\",\r\n    \"username\": \"Nicky\",\r\n    \"password\": \"123\",\r\n    \"email\": \"nicky@gmail.com\",\r\n    \"phoneNumber\": \"987654321\",\r\n    \"photoProfile\": \"jpg\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://localhost:4000/api/v1/users/profile/2391a99d-a805-4247-809e-ffcc65aff7ad",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "4000",
							"path": [
								"api",
								"v1",
								"users",
								"profile",
								"2391a99d-a805-4247-809e-ffcc65aff7ad"
							]
						}
					},
					"response": []
				},
				{
					"name": "Update Attitude Points",
					"request": {
						"auth": {
							"type": "bearer",
							"bearer": [
								{
									"key": "token",
									"value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJEZXNpIiwiZXhwIjoxNzIxMDY5NTM4LCJpYXQiOjE3MjA3MDk1MzgsImlkIjoiNTBhOTU5NDMtZGMxMi00YzFiLTg2ZDYtYTg1ZWQwNjRjNzI0IiwidXNlcm5hbWUiOiJEZXNpIn0.cDK10v4bU9zSH-A6dA6r7JCh0Bx7qnQvfPPAVtP95pI",
									"type": "string"
								}
							]
						},
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"attitudePoints\": 5\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://localhost:4000/api/v1/users/seekers/attitude/e628d076-cc31-44a1-8325-b6ed71479be8",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "4000",
							"path": [
								"api",
								"v1",
								"users",
								"seekers",
								"attitude",
								"e628d076-cc31-44a1-8325-b6ed71479be8"
							]
						}
					},
					"response": []
				}
			]
		},
		{
			"name": "Seeker",
			"item": [
				{
					"name": "Register Seekers",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"username\": \"Desi p\",\r\n    \"password\": \"123\",\r\n    \"fullname\": \"Desi Permatasari\",\r\n    \"email\": \"desi@gmail.com\",\r\n    \"phoneNumber\": \"123456789\",\r\n    \"atitudePoints\": 0,\r\n    \"status\":\"admin\",\r\n    \"photoProfile\": \"jpg\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://localhost:4000/api/v1/seekers/register",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "4000",
							"path": [
								"api",
								"v1",
								"seekers",
								"register"
							]
						}
					},
					"response": []
				},
				{
					"name": "Login Seekers",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"username\": \"Desi p\",\r\n    \"password\": \"123\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://localhost:4000/api/v1/seekers/login",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "4000",
							"path": [
								"api",
								"v1",
								"seekers",
								"login"
							]
						}
					},
					"response": []
				},
				{
					"name": "Get Seeker By Id",
					"request": {
						"auth": {
							"type": "bearer",
							"bearer": [
								{
									"key": "token",
									"value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJEZXNpIiwiZXhwIjoxNzIxMDcxMjcyLCJpYXQiOjE3MjA3MTEyNzIsImlkIjoiZTYyOGQwNzYtY2MzMS00NGExLTgzMjUtYjZlZDcxNDc5YmU4IiwidXNlcm5hbWUiOiJEZXNpIn0.SXUPv0uUaWEoRKVc7AVuaqz0SaT_Fyn77d7CyEJ76v8",
									"type": "string"
								}
							]
						},
						"method": "GET",
						"header": [],
						"url": {
							"raw": "http://localhost:4000/api/v1/seekers/profile/e628d076-cc31-44a1-8325-b6ed71479be8",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "4000",
							"path": [
								"api",
								"v1",
								"seekers",
								"profile",
								"e628d076-cc31-44a1-8325-b6ed71479be8"
							]
						}
					},
					"response": []
				},
				{
					"name": "Get All Seekers",
					"request": {
						"auth": {
							"type": "bearer",
							"bearer": [
								{
									"key": "token",
									"value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJEZXNpIiwiZXhwIjoxNzIxMDcxMjcyLCJpYXQiOjE3MjA3MTEyNzIsImlkIjoiZTYyOGQwNzYtY2MzMS00NGExLTgzMjUtYjZlZDcxNDc5YmU4IiwidXNlcm5hbWUiOiJEZXNpIn0.SXUPv0uUaWEoRKVc7AVuaqz0SaT_Fyn77d7CyEJ76v8",
									"type": "string"
								}
							]
						},
						"method": "GET",
						"header": [],
						"url": {
							"raw": "http://localhost:4000/api/v1/seekers/profile/getall",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "4000",
							"path": [
								"api",
								"v1",
								"seekers",
								"profile",
								"getall"
							]
						}
					},
					"response": []
				},
				{
					"name": "Update Profile",
					"request": {
						"auth": {
							"type": "bearer",
							"bearer": [
								{
									"key": "token",
									"value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJEZXNpIiwiZXhwIjoxNzIxMDcyMTgxLCJpYXQiOjE3MjA3MTIxODEsImlkIjoiZTYyOGQwNzYtY2MzMS00NGExLTgzMjUtYjZlZDcxNDc5YmU4IiwidXNlcm5hbWUiOiJEZXNpIn0.sDCEBHBQSFCIgfl__FphZAmpAIXuUscNx6lRwuzrsak",
									"type": "string"
								}
							]
						},
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"username\": \"Desi\",\r\n    \"password\": \"123\",\r\n    \"fullname\": \"Desi PermataSari\",\r\n    \"email\": \"desi@gmail.com\",\r\n    \"phoneNumber\": \"123456789\",\r\n    \"atitudePoints\": 0,\r\n    \"status\":\"admin\",\r\n    \"photoProfile\": \"jpg\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://localhost:4000/api/v1/seekers/profile/e628d076-cc31-44a1-8325-b6ed71479be8",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "4000",
							"path": [
								"api",
								"v1",
								"seekers",
								"profile",
								"e628d076-cc31-44a1-8325-b6ed71479be8"
							]
						}
					},
					"response": []
				},
				{
					"name": "Delete Seeker",
					"request": {
						"auth": {
							"type": "bearer",
							"bearer": [
								{
									"key": "token",
									"value": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJEZXNpIiwiZXhwIjoxNzIxMDcyNDM0LCJpYXQiOjE3MjA3MTI0MzQsImlkIjoiMzI3OWQ1MjQtZTYxZi00YWZiLWIzYjEtOWE3MTkwNDhhZDZmIiwidXNlcm5hbWUiOiJEZXNpIHAifQ.JO7LjjxEi9c2xbHU9w_MpVIx5VJJjoSVa84YE1EEfWY",
									"type": "string"
								}
							]
						},
						"method": "DELETE",
						"header": [],
						"url": {
							"raw": "http://localhost:4000/api/v1/seekers/profile/3279d524-e61f-4afb-b3b1-9a719048ad6f",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "4000",
							"path": [
								"api",
								"v1",
								"seekers",
								"profile",
								"3279d524-e61f-4afb-b3b1-9a719048ad6f"
							]
						}
					},
					"response": []
				}
			]
		},
		{
			"name": "Transaction",
			"item": [
				{
					"name": "Create Transaction",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"roomId\":\"c23f20e6-f94b-418c-ba54-b06d014a742f\",\r\n    \"seekerId\":\"e628d076-cc31-44a1-8325-b6ed71479be8\",\r\n    \"startDate\": \"2024-08-01\",\r\n    \"voucherId\": \"\",\r\n    \"months\": 2,\r\n    \"payLater\": true,\r\n    \"dueDate\": \"2024-10-01\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://localhost:4000/api/v1/trans/create",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "4000",
							"path": [
								"api",
								"v1",
								"trans",
								"create"
							]
						}
					},
					"response": []
				},
				{
					"name": "Get Transaction By ID",
					"request": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "http://localhost:4000/api/v1/trans/e7a5b7ed-c942-4b21-877b-79fbf9d007b8",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "4000",
							"path": [
								"api",
								"v1",
								"trans",
								"e7a5b7ed-c942-4b21-877b-79fbf9d007b8"
							]
						}
					},
					"response": []
				},
				{
					"name": "Get List Transaction",
					"request": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "http://localhost:4000/api/v1/trans/list",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "4000",
							"path": [
								"api",
								"v1",
								"trans",
								"list"
							]
						}
					},
					"response": []
				},
				{
					"name": "Get List PayLater Transaction",
					"request": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "http://localhost:4000/api/v1/trans/paylater/list",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "4000",
							"path": [
								"api",
								"v1",
								"trans",
								"paylater",
								"list"
							]
						}
					},
					"response": []
				},
				{
					"name": "Update PayLater Transaction",
					"request": {
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"transId\": \"c23a8b06-911a-48dc-99d6-0cc8e5b1823e\",\r\n    \"seekerId\": \"e628d076-cc31-44a1-8325-b6ed71479be8\",\r\n    \"total\": 4000000\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://localhost:4000/api/v1/trans/paylater",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "4000",
							"path": [
								"api",
								"v1",
								"trans",
								"paylater"
							]
						}
					},
					"response": []
				},
				{
					"name": "Acc Payment",
					"request": {
						"method": "PUT",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"transId\": \"e7a5b7ed-c942-4b21-877b-79fbf9d007b8\",\r\n    \"paymentStatus\": \"success\"\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://localhost:4000/api/v1/trans/payment",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "4000",
							"path": [
								"api",
								"v1",
								"trans",
								"payment"
							]
						}
					},
					"response": []
				}
			]
		},
		{
			"name": "Voucher",
			"item": [
				{
					"name": "Create Voucher",
					"request": {
						"method": "POST",
						"header": [],
						"body": {
							"mode": "raw",
							"raw": "{\r\n    \"name\": \"Voucher Desi\",\r\n    \"expiredDate\": \"2024-11-01\",\r\n    \"seekerId\": \"e628d076-cc31-44a1-8325-b6ed71479be8\",\r\n    \"percentAmount\": 30\r\n}",
							"options": {
								"raw": {
									"language": "json"
								}
							}
						},
						"url": {
							"raw": "http://localhost:4000/api/v1/voucher/create",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "4000",
							"path": [
								"api",
								"v1",
								"voucher",
								"create"
							]
						}
					},
					"response": []
				},
				{
					"name": "Get All Vouchers",
					"request": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "http://localhost:4000/api/v1/voucher",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "4000",
							"path": [
								"api",
								"v1",
								"voucher"
							]
						}
					},
					"response": []
				},
				{
					"name": "Get Voucher By ID",
					"request": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "http://localhost:4000/api/v1/voucher/28a2d5a5-4cf7-4e37-9546-af1aadc31be5",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "4000",
							"path": [
								"api",
								"v1",
								"voucher",
								"28a2d5a5-4cf7-4e37-9546-af1aadc31be5"
							]
						}
					},
					"response": []
				},
				{
					"name": "Get Voucher By Seeker ID",
					"request": {
						"method": "GET",
						"header": [],
						"url": {
							"raw": "http://localhost:4000/api/v1/voucher/seeker/e628d076-cc31-44a1-8325-b6ed71479be8",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "4000",
							"path": [
								"api",
								"v1",
								"voucher",
								"seeker",
								"e628d076-cc31-44a1-8325-b6ed71479be8"
							]
						}
					},
					"response": []
				},
				{
					"name": "Delete Voucher",
					"request": {
						"method": "DELETE",
						"header": [],
						"url": {
							"raw": "http://localhost:4000/api/v1/voucher",
							"protocol": "http",
							"host": [
								"localhost"
							],
							"port": "4000",
							"path": [
								"api",
								"v1",
								"voucher"
							]
						}
					},
					"response": []
				}
			]
		}
	]
}