import React from 'react';
import Image from 'next/image';
import { Post } from '@/types';
import moment from 'moment';
import 'pure-react-carousel/dist/react-carousel.es.css';
import Link from 'next/link';
import Toaster from '@/utils/toaster';
import { Buildings } from '@phosphor-icons/react';
import { CarouselProvider, Slider, Slide, Dot } from 'pure-react-carousel';
import 'pure-react-carousel/dist/react-carousel.es.css';
import renderContentWithLinks from '@/utils/render_content_with_links';
import { USER_PROFILE_PIC_URL, POST_PIC_URL } from '@/config/routes';
import postHandler from '@/handlers/post_handler';
import { SERVER_ERROR } from '@/config/errors';
import { X } from '@phosphor-icons/react/dist/ssr';

interface Props {
  post: Post;
  setPosts: React.Dispatch<React.SetStateAction<Post[]>>;
}

const Post = ({ post, setPosts }: Props) => {
  const removeFlag = () => {
    const toaster = Toaster.startLoad('Removing Flag', post.id);
    const URL = `/flags/posts/${post.id}`;
    postHandler(URL, {})
      .then(res => {
        if (res.statusCode === 200) {
          setPosts(prev => prev.filter(p => p.id != post.id));
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
          <Link href={`${`/explore/user/${post.user.username}`}`} className="rounded-full">
            <Image
              crossOrigin="anonymous"
              width={100}
              height={100}
              alt={'User Pic'}
              src={`${USER_PROFILE_PIC_URL}/${post.user.profilePic}`}
              className={'rounded-full w-8 h-8'}
            />
          </Link>
        </div>
        <div className="w-[calc(100%-32px)] flex flex-col gap-1">
          <div className="w-full h-fit flex justify-between">
            <Link href={`${`/explore/user/${post.user.username}`}`} className="font-medium flex items-center gap-1">
              {post.user.name}
              {post.user.isOrganization ? <Buildings weight="duotone" /> : <></>}
              <div className="text-xs font-normal text-gray-500">@{post.user.username}</div>
            </Link>
            <div className="flex-center gap-2 text-xs text-gray-400">
              {post.isEdited && <div>(edited)</div>}
              <div>{moment(post.postedAt).fromNow()}</div>
            </div>
          </div>
          {post.images && post.images.length > 0 && (
            <CarouselProvider
              naturalSlideHeight={580}
              naturalSlideWidth={1000}
              totalSlides={post.images.length}
              visibleSlides={1}
              infinite={true}
              dragEnabled={post.images.length != 1}
              touchEnabled={post.images.length != 1}
              isPlaying={false}
              className="w-full rounded-lg flex flex-col items-center justify-center relative"
            >
              <Slider className={`w-full rounded-lg`}>
                {post.images.map((image, index) => {
                  return (
                    <Slide
                      index={index}
                      key={index}
                      className={`w-full rounded-lg flex items-center justify-center gap-2`}
                    >
                      <Image
                        crossOrigin="anonymous"
                        width={500}
                        height={500}
                        className="w-full"
                        alt={'Post Pic'}
                        src={`${POST_PIC_URL}/${image}`}
                        placeholder="blur"
                        blurDataURL={(post.hashes && post.hashes[index]) || 'no-hash'}
                      />
                    </Slide>
                  );
                })}
              </Slider>
              <div className={`${post.images.length === 1 ? 'hidden' : ''} absolute bottom-5`}>
                {post.images.map((_, i) => {
                  return <Dot key={i} slide={i} />;
                })}
              </div>
            </CarouselProvider>
          )}
          <div className="w-full text-sm  whitespace-pre-wrap mb-2">
            {renderContentWithLinks(post.content, post.taggedUsers)}
          </div>
        </div>
      </div>
      <X onClick={removeFlag} className="w-full bg-red-100 rounded-b-lg cursor-pointer h-4" />
    </div>
  );
};

export default Post;
