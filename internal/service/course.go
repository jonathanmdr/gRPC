package service

import (
	"context"
	"io"

	"github.com/jonathanmdr/gRPC/internal/database"
	"github.com/jonathanmdr/gRPC/internal/pb"
)

type CourseService struct {
	pb.UnimplementedCourseServiceServer
	CourseDB   database.Course
	CategoryDB database.Category
}

func NewCourseService(courseDb database.Course, categoryDb database.Category) *CourseService {
	return &CourseService{
		CourseDB:   courseDb,
		CategoryDB: categoryDb,
	}
}

func (c *CourseService) GetCourse(ctx context.Context, in *pb.GetCourseRequest) (*pb.CourseResponse, error) {
	course, err := c.CourseDB.FindById(in.Id)
	if err != nil {
		return nil, err
	}
	category, err := c.CategoryDB.FindById(course.CategoryID)
	if err != nil {
		return nil, err
	}
	return &pb.CourseResponse{
		Id:          course.ID,
		Name:        course.Name,
		Description: course.Description,
		Category: &pb.CategoryResponse{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		},
	}, nil
}

func (c *CourseService) GetCourses(context.Context, *pb.EmptyRequest) (*pb.CoursesResponse, error) {
	courses, err := c.CourseDB.FindAll()
	if err != nil {
		return nil, err
	}
	var coursesResponse []*pb.CourseResponse
	for _, course := range courses {
		category, err := c.CategoryDB.FindById(course.CategoryID)
		if err != nil {
			return nil, err
		}
		var courseResponse = &pb.CourseResponse{
			Id:          course.ID,
			Name:        course.Name,
			Description: course.Description,
			Category: &pb.CategoryResponse{
				Id:          category.ID,
				Name:        category.Name,
				Description: category.Description,
			},
		}
		coursesResponse = append(coursesResponse, courseResponse)
	}
	return &pb.CoursesResponse{
		Courses: coursesResponse,
	}, nil
}

func (c *CourseService) CreateCourse(ctx context.Context, in *pb.CreateCourseRequest) (*pb.CourseResponse, error) {
	category, err := c.CategoryDB.FindById(in.CategoryId)
	if err != nil {
		return nil, err
	}
	course, err := c.CourseDB.Create(in.Name, in.Description, in.CategoryId)
	if err != nil {
		return nil, err
	}
	return &pb.CourseResponse{
		Id:          course.ID,
		Name:        course.Name,
		Description: course.Description,
		Category: &pb.CategoryResponse{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		},
	}, nil
}

func (c *CourseService) CreateCourseStream(stream pb.CourseService_CreateCourseStreamServer) error {
	courses := &pb.CoursesResponse{}

	for {
		courseRequest, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(courses)
		}
		if err != nil {
			return err
		}
		category, err := c.CategoryDB.FindById(courseRequest.CategoryId)
		if err != nil {
			return err
		}
		course, err := c.CourseDB.Create(
			courseRequest.Name,
			courseRequest.Description,
			courseRequest.CategoryId,
		)
		if err != nil {
			return err
		}
		var courseResponse = &pb.CourseResponse{
			Id:          course.ID,
			Name:        course.Name,
			Description: course.Description,
			Category: &pb.CategoryResponse{
				Id:          category.ID,
				Name:        category.Name,
				Description: category.Description,
			},
		}
		courses.Courses = append(courses.Courses, courseResponse)
	}
}

func (c *CourseService) CreateCourseStreamBidirectional(stream pb.CourseService_CreateCourseStreamBidirectionalServer) error {
	for {
		courseRequest, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		category, err := c.CategoryDB.FindById(courseRequest.CategoryId)
		if err != nil {
			return err
		}
		course, err := c.CourseDB.Create(
			courseRequest.Name,
			courseRequest.Description,
			courseRequest.CategoryId,
		)
		if err != nil {
			return err
		}
		var courseResponse = &pb.CourseResponse{
			Id:          course.ID,
			Name:        course.Name,
			Description: course.Description,
			Category: &pb.CategoryResponse{
				Id:          category.ID,
				Name:        category.Name,
				Description: category.Description,
			},
		}
		err = stream.Send(courseResponse)
		if err != nil {
			return err
		}
	}
}
