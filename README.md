# Go Toolkit 🛠️ 

[![Go Version](https://img.shields.io/badge/Go-1.21%2B-blue?logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)
[![Build Status](https://img.shields.io/github/actions/workflow/status/yourusername/go-collections/test.yml?logo=github)](https://github.com/yourusername/go-collections/actions)
[![Coverage](https://img.shields.io/badge/Coverage-95%25-brightgreen)](https://gocover.io/github.com/yourusername/go-collections)
[![GoDoc](https://pkg.go.dev/badge/github.com/yourusername/go-collections)](https://pkg.go.dev/github.com/yourusername/go-collections)

一个现代化易上手的Go语言泛型工具库，提供类似Java风格的实用操作方法，同时保持Go语言的高性能特性。✨

## 功能特性

### 当前支持的功能 ✅

#### **slicex 功能分组表**
| 类型        | 功能分组     | 已实现方法                                                                                     | 版本     |
|-----------|----------|-------------------------------------------------------------------------------------------|--------|
| **slice** | **索引查询** | `IndexOf`、`IndexOfFunc`、`LastIndexOf`、`LastIndexOfFunc`、`IndexOfAll`、`IndexOfAllFunc`     | v0.0.1 |
|           | **包含查询** | `Contains`、`ContainsFunc`、`ContainsAll`、`ContainsAllFunc`、`ContainsAny`、`ContainsAnyFunc` |        |
|           | **条件查询** | `Find`、`FindAll`                                                                          |        |
|           | **聚合计算** | `Max`、`Min`、`Sum`                                                                         |        |
|           | **操作类**  | `Reverse`、`ReverseSelf`                                                                   |        |

#### **stringx 功能分组表**
| 类型         | 功能分组     | 已实现方法                                                                 | 版本     |
|------------|----------|-----------------------------------------------------------------------|--------|
| **string** | **操作类** | `IsEmpty`、`IsNotEmpty`、`IsBlank`、`IsNotBlank` | v0.0.1 |

#### **randx 功能分组表**
| 类型         | 功能分组     | 已实现方法                                     | 版本     |
|------------|----------|-------------------------------------------|--------|
| **random** | **操作类** | `Int` | v0.0.1 |

#### **logx 功能分组表**
| 类型      | 功能分组     | 已实现方法                         | 版本     |
|---------|----------|-------------------------------|--------|
| **log** | **操作类** | `Debug`、`Info`、`Warn`、`Error` | v0.0.1 |

### 当前支持的数据结构 ✅
#### **stack 数据结构表**
| 数据结构      | 数据结构分组          | 已实现方法                                                | 版本     |
|-----------|-----------------|------------------------------------------------------|--------|
| **stack** | **array_stack** | `Push`、`Pop`、`Peek`、`Empty` | v0.0.1 | |
|           | **list_stack**  | `Push`、`Pop`、`Peek`、`Empty` |        | |

#### **queue 数据结构表**
| 数据结构      | 数据结构分组                   | 已实现方法                                  | 版本     |
|-----------|--------------------------|----------------------------------------|--------|
| **queue** | **array_queue**          | `Offer`、`Poll`、`Peek`、`IsEmpty`、`Size` | v0.0.1 | |
|           | **linked_queue**         | `Offer`、`Poll`、`Peek`、`IsEmpty`、`Size` |        | |
|           | **priority_queue**       | `Offer`、`Poll`、`Peek`、`IsEmpty`、`Size` |        | |
|           | **array_blocking_queue** | `Put`、`Take`                           |        | |

#### **bitmap 数据结构表**
| 数据结构       | 数据结构分组 | 已实现方法                             | 版本     |
|------------|--------|-----------------------------------|--------|
| **bitmap** |        | `Set`、`Contains`、`Reset`、`Size`   | v0.0.1 | |

#### **pair 数据结构表**
| 数据结构     | 数据结构分组 | 已实现方法            | 版本     |
|----------|--------|------------------|--------|
| **pair** |        | `String`、`Split` | v0.0.1 | |

### 开发计划中 🚧

| 数据结构            | 计划方法 | 状态    |
|-----------------|------|---------|
| **Set**         | 待定   | 计划中 📅|
| **UnionSet**    | 待定   | 计划中 📅|
| **Tree**        | 待定   | 计划中 📅|
| **HashMap**     | 待定   | 计划中 📅|
| **BloomFilter** | 待定   | 计划中 📅|
| **Beanx**       | 待定   | 计划中 📅|
| **ThreadPool**  | 待定   | 计划中 📅|
| **Retry**       | 待定   | 计划中 📅|
| **Sync**        | 待定   | 计划中 📅|

## 📦 安装

```bash
go get github.com/ricejson/gotool
```
