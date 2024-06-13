import React from 'react';
import Image from 'next/image';
import { Project } from '@/types';
import 'pure-react-carousel/dist/react-carousel.es.css';
import Link from 'next/link';
import Toaster from '@/utils/toaster';
import { Buildings } from '@phosphor-icons/react';
import 'pure-react-carousel/dist/react-carousel.es.css';
import { PROJECT_PIC_URL } from '@/config/routes';
import postHandler from '@/handlers/post_handler';
import { SERVER_ERROR } from '@/config/errors';
import { X } from '@phosphor-icons/react/dist/ssr';

interface Props {
  project: Project;
  setProjects: React.Dispatch<React.SetStateAction<Project[]>>;
}

const ProjectComponent = ({ project, setProjects }: Props) => {
  const removeFlag = () => {
    const toaster = Toaster.startLoad('Removing Flag', project.id);
    const URL = `/flags/projects/${project.id}`;
    postHandler(URL, {})
      .then(res => {
        if (res.statusCode === 200) {
          setProjects(prev => prev.filter(p => p.id != project.id));
          Toaster.stopLoad(toaster, 'Flag Removed', 1);
        } else {
          if (res.data.message) Toaster.stopLoad(toaster, res.data.message, 0);
          else {
            Toaster.stopLoad(toaster, SERVER_ERROR, 0);
          }
        }
      })
      .catch(err => {
        Toaster.stopLoad(toaster, SERVER_ERROR, 0);
      });
  };
  return (
    <div className="w-full">
      <div className="w-full relative bg-white font-primary flex gap-1 rounded-lg rounded-b-none p-2 text-primary_black border-gray-300 border-b-0 border-[1px]">
        <div className="h-full">
          <Link href={`${`/explore/user/${project.user.username}`}`} className="rounded-full">
            <Image
              crossOrigin="anonymous"
              width={100}
              height={100}
              alt={'Project Pic'}
              src={`${PROJECT_PIC_URL}/${project.user.profilePic}`}
              className={'rounded-full w-8 h-8'}
            />
          </Link>
        </div>
        <div className="w-[calc(100%-32px)] flex flex-col gap-1">
          <div className="w-full h-fit flex justify-between">
            <Link href={`${`/explore/user/${project.user.username}`}`} className="font-medium flex items-center gap-1">
              {project.user.name}
              {project.user.isOrganization ? <Buildings weight="duotone" /> : <></>}
              <div className="text-xs font-normal text-gray-500">@{project.user.username}</div>
              <div className="text-xs font-normal text-gray-500">@{project.title}</div>
              <div className="text-xs font-normal text-gray-500">@{project.category}</div>
              <div className="text-xs font-normal text-gray-500">@{project.description}</div>
              <div className="text-xs font-normal text-gray-500">@{project.createdAt.toLocaleDateString()}</div>
            </Link>
          </div>
        </div>
      </div>
      <X onClick={removeFlag} className="w-full bg-red-100 rounded-b-lg cursor-pointer h-4" />
    </div>
  );
};

export default ProjectComponent;